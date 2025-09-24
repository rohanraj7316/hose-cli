package cmd

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var lineLog = log.New(os.Stderr, "", 0)

var RootCmd = &cobra.Command{
	Use:   "hose",
	Short: "Hose is a CLI for managing the ASL application",
}

func init() {
	// Silence Cobra's default usage and error printing for snappier output
	RootCmd.SilenceUsage = true
	RootCmd.SilenceErrors = true
	RootCmd.AddCommand(encryptionCmd)
	RootCmd.AddCommand(decryptionCmd)
}

func getFlagString(cmd *cobra.Command, name string) (string, error) {
	return cmd.Flags().GetString(name)
}

func logError(message string, err error) {
	// Minimal, single-line error log to stderr
	w := lineLog.Writer()
	_, _ = io.WriteString(w, message+": ")
	_, _ = io.WriteString(w, err.Error())
	_, _ = io.WriteString(w, "\n")
}
