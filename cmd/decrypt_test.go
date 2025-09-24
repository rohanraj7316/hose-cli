package cmd

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/rohanraj7316/hose-cli/utils/hose"
)

func TestDecryptCommand_Success_JSON(t *testing.T) {
	resetCobraState(t)
	original := "hello world"
	secretKey := strings.Repeat("b", 32)
	publicKey := generatePublicKeyB64(t)

	_, encryptedPayload, err := hose.New().Encrypt(original, secretKey, publicKey)
	if err != nil {
		t.Fatalf("failed to prepare encrypted payload: %v", err)
	}

	RootCmd.SetArgs([]string{
		"decrypt",
		"--encrypted-payload", encryptedPayload,
		"--secret-key", secretKey,
		"--json",
	})

	out := captureStdout(t, func() {
		if err := RootCmd.Execute(); err != nil {
			t.Fatalf("execute failed: %v", err)
		}
	})

	var got DecryptCmdInput
	if err := json.Unmarshal([]byte(strings.TrimSpace(out)), &got); err != nil {
		t.Fatalf("invalid JSON output: %v; out=%s", err, out)
	}
	if got.DecryptedPayload != original {
		t.Fatalf("unexpected decrypted payload: got %q want %q", got.DecryptedPayload, original)
	}
	if got.Latency == "" {
		t.Fatalf("expected latency_us to be present")
	}
}

func TestDecryptCommand_MissingFlags(t *testing.T) {
	resetCobraState(t)
	RootCmd.SetArgs([]string{"decrypt", "--encrypted-payload", "abc.def.ghi"})
	if err := RootCmd.Execute(); err == nil {
		t.Fatalf("expected error due to missing required flags, got nil")
	}
}

func TestDecryptCommand_InvalidInput(t *testing.T) {
	resetCobraState(t)
	// Invalid secret key length should cause decryption to fail and be logged.
	RootCmd.SetArgs([]string{
		"decrypt",
		"--encrypted-payload", "YW5vbmNl.YXV0aHRhZw==.cGF5bG9hZA==", // nonce.authTag.payload (base64) placeholder
		"--secret-key", "short",
	})

	logs := captureLoggerOutput(t, func() {
		if err := RootCmd.Execute(); err != nil {
			t.Fatalf("unexpected execute error (Run handles errors internally): %v", err)
		}
	})

	if !strings.Contains(logs, "Error decrypting payload") {
		t.Fatalf("expected log about decryption error, got: %s", logs)
	}
}
