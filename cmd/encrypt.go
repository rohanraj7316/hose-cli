package cmd

import (
	"encoding/json"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/rohanraj7316/hose-cli/utils/hose"
	"github.com/spf13/cobra"
)

type EncryptCmdOutput struct {
	ApiEncryptionKey string `json:"api_encryption_key"`
	EncryptedPayload string `json:"encrypted_payload"`
	Latency          string `json:"latency"`
}

var encryptionCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt is a CLI for encrypting the data",
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		payload, err := getFlagString(cmd, "payload")
		if err != nil {
			logError("Error getting payload flag", err)
			return
		}

		secretKey, err := getFlagString(cmd, "secret-key")
		if err != nil {
			logError("Error getting secret key flag", err)
			return
		}

		publicKey, err := getFlagString(cmd, "public-key")
		if err != nil {
			logError("Error getting public key flag", err)
			return
		}

		jsonFlag, err := cmd.Flags().GetBool("json")
		if err != nil {
			logError("Error getting json flag", err)
			return
		}

		apiEncryptionKey, encryptedPayload, err := hose.New().Encrypt(payload, secretKey, publicKey)
		if err != nil {
			logError("Error encrypting payload", err)
			return
		}

		latencyUs := strconv.FormatInt(time.Since(start).Microseconds(), 10)
		outputResultForEncrypt(apiEncryptionKey, encryptedPayload, jsonFlag, latencyUs)
	},
}

func init() {
	encryptionCmd.Flags().StringP("payload", "p", "", "payload to encrypt")
	encryptionCmd.MarkFlagRequired("payload")

	encryptionCmd.Flags().StringP("secret-key", "s", "", "secret key to encrypt the payload")
	encryptionCmd.MarkFlagRequired("secret-key")

	encryptionCmd.Flags().StringP("public-key", "k", "", "public key to encrypt the payload")
	encryptionCmd.MarkFlagRequired("public-key")

	encryptionCmd.Flags().BoolP("json", "", false, "output in json format")
}

func outputResultForEncrypt(apiEncryptionKey, encryptedPayload string, jsonFlag bool, latencyUs string) {
	if jsonFlag {
		enc := json.NewEncoder(os.Stdout)
		_ = enc.Encode(EncryptCmdOutput{ApiEncryptionKey: apiEncryptionKey, EncryptedPayload: encryptedPayload, Latency: latencyUs})
	} else {
		io.WriteString(os.Stdout, "Encrypted Payload: "+" "+encryptedPayload+"\n")
		io.WriteString(os.Stdout, "API Encryption Key: "+" "+apiEncryptionKey+"\n")
	}
}
