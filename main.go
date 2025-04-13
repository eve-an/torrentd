package main

import (
	"os"

	"github.com/eve-an/torrentd/cmd"
	"github.com/eve-an/torrentd/pkg/status"
	"github.com/spf13/cobra"
)

func main() {
	service := status.NewStatusService()

	rootCmd := &cobra.Command{Use: "torrentd"}
	rootCmd.AddCommand(cmd.NewStatusCommand(service))

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
