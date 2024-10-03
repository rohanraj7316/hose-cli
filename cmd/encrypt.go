package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/rohanraj7316/hose-cli/utils/hose"
	"github.com/spf13/cobra"
)

type EncryptCmdOutput struct {
	ApiEncryptionKey string `json:"api_encryption_key"`
	EncryptedPayload string `json:"encrypted_payload"`
}

var encryptionCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt is a CLI for encrypting the data",
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := cmd.Flags().GetString("payload")
		if err != nil {
			fmt.Println("Error getting payload flag:", err)
			return
		}

		secretKey, err := cmd.Flags().GetString("secret-key")
		if err != nil {
			fmt.Println("Error getting secret key flag:", err)
			return
		}

		publicKey, err := cmd.Flags().GetString("public-key")
		if err != nil {
			fmt.Println("Error getting public key flag:", err)
			return
		}

		jsonFlag, err := cmd.Flags().GetBool("json")
		if err != nil {
			fmt.Println("Error getting json flag:", err)
			return
		}

		apiEncryptionKey, encryptedPayload, err := hose.New().Encrypt(payload, secretKey, publicKey)
		if err != nil {
			fmt.Println("Error encrypting payload:", err)
			return
		}

		if jsonFlag {
			output := EncryptCmdOutput{
				ApiEncryptionKey: apiEncryptionKey,
				EncryptedPayload: encryptedPayload,
			}

			jsonOutput, err := json.Marshal(output)
			if err != nil {
				fmt.Println("Error converting output to JSON:", err)
				return
			}
			fmt.Println(string(jsonOutput))
		} else {
			fmt.Println("Encrypted Payload: ", encryptedPayload)
			fmt.Println("API Encryption Key: ", apiEncryptionKey)
		}
	},
}

func init() {
	encryptionCmd.PersistentFlags().StringP("payload", "p", "", "payload to encrypt")
	encryptionCmd.MarkPersistentFlagRequired("payload")

	encryptionCmd.PersistentFlags().StringP("secret-key", "s", "", "secret key to encrypt the payload")
	encryptionCmd.MarkPersistentFlagRequired("secret-key")

	encryptionCmd.PersistentFlags().StringP("public-key", "k", "", "public key to encrypt the payload")
	encryptionCmd.MarkPersistentFlagRequired("public-key")

	encryptionCmd.PersistentFlags().BoolP("json", "", false, "output in json format")
}
