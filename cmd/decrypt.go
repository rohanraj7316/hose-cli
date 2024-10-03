package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/rohanraj7316/hose-cli/utils/hose"
	"github.com/spf13/cobra"
)

type DecryptCmdInput struct {
	DecryptedPayload string `json:"decrypted_payload"`
}

var decryptionCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt is a CLI for decrypting the data",
	Run: func(cmd *cobra.Command, args []string) {
		encryptedPayload, err := cmd.Flags().GetString("encrypted-payload")
		if err != nil {
			fmt.Println("Error getting encrypted payload flag:", err)
			return
		}

		secretKey, err := cmd.Flags().GetString("secret-key")
		if err != nil {
			fmt.Println("Error getting secret key flag:", err)
			return
		}

		jsonFlag, err := cmd.Flags().GetBool("json")
		if err != nil {
			fmt.Println("Error getting json flag:", err)
			return
		}

		decryptedPayload, err := hose.New().Decrypt(encryptedPayload, secretKey)
		if err != nil {
			fmt.Println("Error decrypting payload:", err)
			return
		}

		if jsonFlag {
			output := DecryptCmdInput{
				DecryptedPayload: decryptedPayload,
			}

			jsonOutput, err := json.Marshal(output)
			if err != nil {
				fmt.Println("Error converting output to JSON:", err)
				return
			}
			fmt.Println(string(jsonOutput))
		} else {
			fmt.Println("Decrypted Payload: ", decryptedPayload)
		}
	},
}

func init() {
	decryptionCmd.PersistentFlags().StringP("encrypted-payload", "e", "", "encrypted payload to decrypt")
	decryptionCmd.MarkPersistentFlagRequired("encrypted-payload")

	decryptionCmd.PersistentFlags().StringP("secret-key", "s", "", "secret key to decrypt the payload")
	decryptionCmd.MarkPersistentFlagRequired("secret-key")

	decryptionCmd.PersistentFlags().BoolP("json", "", false, "output in json format")
}
