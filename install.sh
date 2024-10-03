#!/bin/bash

# Function to check if a command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Function to install the CLI
install_cli() {
    echo "Installing hose CLI..."

    # Check if Go is installed
    if ! command_exists go; then
        echo "Go is not installed. Please install Go and try again."
        exit 1
    fi

    # Clone the repository
    if [ -d "hose" ]; then
        rm -rf hose
    fi
    git clone https://github.com/rohanraj7316/hose-cli.git

    # Navigate to the project directory
    cd hose-cli || exit

    # Build the CLI
    go build -o hose

    # Move the binary to /usr/local/bin
    sudo mv hose /usr/local/bin/

    # Verify installation
    if command_exists hose; then
        echo "hose CLI installed successfully!"
    else
        echo "Failed to install hose CLI."
        exit 1
    fi
}

# Detect the OS
OS="$(uname -s)"
case "${OS}" in
    Linux*)     install_cli;;
    Darwin*)    install_cli;;
    *)          echo "Unsupported OS: ${OS}"; exit 1;;
esac
