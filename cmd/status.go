package cmd

import (
	"fmt"
	"math"
	"os"

	"github.com/eve-an/torrentd/pkg/status"
	"github.com/spf13/cobra"
)

func NewStatusCommand(statusService *status.StatusService) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Prints status of torrent download",
		Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath, torrentFilePath := args[0], args[1]

			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			torrentFile, err := os.Open(torrentFilePath)
			if err != nil {
				return err
			}
			defer torrentFile.Close()

			progress, err := statusService.CheckStatus(file, torrentFile)
			if err != nil {
				return err
			}

			fmt.Printf("Progress %d/%d - %.2f%%\n", progress.CompletedBlocks, progress.TotalBlocks, math.Round(progress.Progress*100))
			return nil
		},
	}
}
