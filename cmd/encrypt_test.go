package cmd

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// generatePublicKeyB64 generates a 2048-bit RSA public key and returns it
// as a base64-encoded DER (PKIX) string as expected by utils/hose.
func generatePublicKeyB64(t *testing.T) string {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("failed generating RSA key: %v", err)
	}
	der, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		t.Fatalf("failed marshalling public key: %v", err)
	}
	return base64.StdEncoding.EncodeToString(der)
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()
	orig := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed creating pipe: %v", err)
	}
	os.Stdout = w
	fn()
	_ = w.Close()
	os.Stdout = orig
	b, _ := io.ReadAll(r)
	return string(b)
}

func captureLoggerOutput(t *testing.T, fn func()) string {
	t.Helper()
	var buf bytes.Buffer
	origWriter := os.Stdout
	lineLog.SetOutput(&buf)
	fn()
	lineLog.SetOutput(origWriter)
	return buf.String()
}

func resetFlagSet(fs *pflag.FlagSet) {
	if fs == nil {
		return
	}
	fs.VisitAll(func(f *pflag.Flag) {
		_ = f.Value.Set(f.DefValue)
		f.Changed = false
	})
}

func resetCommandFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	resetFlagSet(cmd.Flags())
	resetFlagSet(cmd.PersistentFlags())
}

func resetCobraState(t *testing.T) {
	t.Helper()
	RootCmd.SilenceUsage = true
	RootCmd.SilenceErrors = true
	resetCommandFlags(RootCmd)
	resetCommandFlags(encryptionCmd)
	resetCommandFlags(decryptionCmd)
}

func TestEncryptCommand_Success_JSON(t *testing.T) {
	resetCobraState(t)
	payload := "test payload"
	secretKey := strings.Repeat("a", 32)
	publicKey := generatePublicKeyB64(t)

	RootCmd.SetArgs([]string{
		"encrypt",
		"--payload", payload,
		"--secret-key", secretKey,
		"--public-key", publicKey,
		"--json",
	})

	out := captureStdout(t, func() {
		if err := RootCmd.Execute(); err != nil {
			t.Fatalf("execute failed: %v", err)
		}
	})

	var got EncryptCmdOutput
	if err := json.Unmarshal([]byte(strings.TrimSpace(out)), &got); err != nil {
		t.Fatalf("invalid JSON output: %v; out=%s", err, out)
	}
	if got.ApiEncryptionKey == "" {
		t.Fatalf("expected api_encryption_key, got empty")
	}
	if _, err := base64.StdEncoding.DecodeString(got.ApiEncryptionKey); err != nil {
		t.Fatalf("api_encryption_key not base64: %v", err)
	}
	if strings.Count(got.EncryptedPayload, ".") != 2 {
		t.Fatalf("encrypted_payload format invalid, expected three segments: %s", got.EncryptedPayload)
	}
}

func TestEncryptCommand_MissingFlags(t *testing.T) {
	resetCobraState(t)
	RootCmd.SetArgs([]string{"encrypt", "--payload", "only-payload"})
	if err := RootCmd.Execute(); err == nil {
		t.Fatalf("expected error due to missing required flags, got nil")
	}
}

func TestEncryptCommand_InvalidInput(t *testing.T) {
	resetCobraState(t)
	// Invalid secret key length triggers AES validation error.
	payload := "test payload"
	secretKey := "short-key"
	publicKey := generatePublicKeyB64(t)

	RootCmd.SetArgs([]string{
		"encrypt",
		"--payload", payload,
		"--secret-key", secretKey,
		"--public-key", publicKey,
	})

	logs := captureLoggerOutput(t, func() {
		if err := RootCmd.Execute(); err != nil {
			t.Fatalf("unexpected execute error (Run handles errors internally): %v", err)
		}
	})

	if !strings.Contains(logs, "Error encrypting payload") {
		t.Fatalf("expected log about encryption error, got: %s", logs)
	}
}
