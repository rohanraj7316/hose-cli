#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'

# Utilities
command_exists() { command -v "$1" >/dev/null 2>&1; }
die() { echo "Error: $*" >&2; exit 1; }
info() { echo "[hose-install] $*"; }

# Detect OS/ARCH
detect_platform() {
    local uos uarch
    uos="$(uname -s)" || uos=""
    uarch="$(uname -m)" || uarch=""
    case "$uos" in
        Linux) OS=linux;;
        Darwin) OS=darwin;;
        *) die "Unsupported OS: $uos";;
    esac
    case "$uarch" in
        x86_64|amd64) ARCH=amd64;;
        aarch64|arm64) ARCH=arm64;;
        armv7l|armv7) ARCH=armv7;;
        *) ARCH="$uarch"; info "Unknown arch '$uarch', proceeding as-is";;
    esac
}

# Choose install directory (non-root friendly by default)
choose_install_dir() {
    if [ -n "${HOSE_INSTALL_DIR:-}" ]; then
        INSTALL_DIR="$HOSE_INSTALL_DIR"
    elif [ -w "/usr/local/bin" ]; then
        INSTALL_DIR="/usr/local/bin"
    else
        INSTALL_DIR="$HOME/.local/bin"
    fi
    mkdir -p "$INSTALL_DIR"
}

ensure_path_note() {
    case ":${PATH}:" in
        *:"${INSTALL_DIR}":*) ;;
        *) info "Note: add '${INSTALL_DIR}' to your PATH to use 'hose' globally";;
    esac
}

# Try prebuilt binary from GitHub Releases if available
try_prebuilt() {
    local url tmp tgz
    url="https://github.com/rohanraj7316/hose-cli/releases/latest/download/hose_${OS}_${ARCH}.tar.gz"
    tmp="$(mktemp -d)"
    tgz="$tmp/hose.tgz"
    trap 'rm -rf "$tmp"' RETURN

    if command_exists curl; then
        if ! curl -fsSL "$url" -o "$tgz"; then
            return 1
        fi
    elif command_exists wget; then
        if ! wget -qO "$tgz" "$url"; then
            return 1
        fi
    else
        return 1
    fi

    if ! tar -xzf "$tgz" -C "$tmp" 2>/dev/null; then
        return 1
    fi
    if [ ! -f "$tmp/hose" ]; then
        # Some archives may contain nested directory
        if ! HOSE_BIN="$(find "$tmp" -type f -name hose | head -n1)"; then
            return 1
        fi
    else
        HOSE_BIN="$tmp/hose"
    fi
    install_binary "$HOSE_BIN"
}

install_binary() {
    local src="$1" dst
    dst="$INSTALL_DIR/hose"
    if mv "$src" "$dst" 2>/dev/null; then
        :
    else
        if command_exists sudo; then
            sudo mv "$src" "$dst"
        else
            die "Permission denied moving binary to '$dst'. Set HOSE_INSTALL_DIR to a writable dir or run with sudo."
        fi
    fi
    chmod +x "$dst"
}

build_from_source() {
    command_exists git || die "git is required to build from source"
    command_exists go || die "Go is not installed. Please install Go (>=1.21) and try again."

    local tmp repo
    tmp="$(mktemp -d)"
    trap 'rm -rf "$tmp"' RETURN
    info "Cloning repository..."
    git clone --depth 1 https://github.com/rohanraj7316/hose-cli.git "$tmp/hose-cli"
    cd "$tmp/hose-cli"
    info "Building..."
    GOFLAGS="-trimpath" go build -ldflags="-s -w" -o hose
    install_binary "./hose"
}

verify_install() {
    if command_exists hose; then
        info "hose CLI installed successfully at: $(command -v hose)"
    else
        info "hose was installed to '$INSTALL_DIR/hose' but is not in PATH"
        ensure_path_note
        exit 1
    fi
}

main() {
    info "Installing hose CLI..."
    detect_platform
    choose_install_dir

    # Silence errors for prebuilt attempt; fallback to source
    if ! try_prebuilt; then
        info "Prebuilt for ${OS}/${ARCH} not available; building from source..."
        build_from_source
    fi

    verify_install
}

main "$@"
