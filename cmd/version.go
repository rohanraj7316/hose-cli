package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hose CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hose CLI v1.0.0")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
