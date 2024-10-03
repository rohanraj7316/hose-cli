package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "hose",
	Short: "Hose is a CLI for managing the ASL application",
}

func init() {
	RootCmd.AddCommand(encryptionCmd)
	RootCmd.AddCommand(decryptionCmd)
}
