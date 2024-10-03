package main

import (
	"context"
	"os"

	"github.com/rohanraj7316/hose/cmd"
)

func main() {
	if err := cmd.RootCmd.ExecuteContext(context.Background()); err != nil {
		os.Exit(1)
	}
}
