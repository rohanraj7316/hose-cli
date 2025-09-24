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

type DecryptCmdInput struct {
	DecryptedPayload string `json:"decrypted_payload"`
	Latency          string `json:"latency"`
}

var decryptionCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt is a CLI for decrypting the data",
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		encryptedPayload, err := getFlagString(cmd, "encrypted-payload")
		if err != nil {
			logError("Error getting encrypted payload flag", err)
			return
		}

		secretKey, err := getFlagString(cmd, "secret-key")
		if err != nil {
			logError("Error getting secret key flag", err)
			return
		}

		jsonFlag, err := cmd.Flags().GetBool("json")
		if err != nil {
			logError("Error getting json flag", err)
			return
		}

		decryptedPayload, err := hose.New().Decrypt(encryptedPayload, secretKey)
		if err != nil {
			logError("Error decrypting payload", err)
			return
		}

		latencyUs := strconv.FormatInt(time.Since(start).Microseconds(), 10)
		outputResultForDecrypt(decryptedPayload, jsonFlag, latencyUs)
	},
}

func init() {
	decryptionCmd.Flags().StringP("encrypted-payload", "e", "", "encrypted payload to decrypt")
	decryptionCmd.MarkFlagRequired("encrypted-payload")

	decryptionCmd.Flags().StringP("secret-key", "s", "", "secret key to decrypt the payload")
	decryptionCmd.MarkFlagRequired("secret-key")

	decryptionCmd.Flags().BoolP("json", "", false, "output in json format")
}

func outputResultForDecrypt(decryptedPayload string, jsonFlag bool, latencyUs string) {
	if jsonFlag {
		enc := json.NewEncoder(os.Stdout)
		_ = enc.Encode(DecryptCmdInput{DecryptedPayload: decryptedPayload, Latency: latencyUs})
	} else {
		io.WriteString(os.Stdout, "Decrypted Payload: "+" "+decryptedPayload+"\n")
	}
}
