package cmd

import (
	"github.com/sid-008/hnCLI/api"
	"github.com/spf13/cobra"
)

var topCmd = &cobra.Command{
	Use:   "top",
	Short: "Get top items",
	Run: func(cmd *cobra.Command, args []string) {
		api.GetTopItems()
	},
}

func init() {
	rootCmd.AddCommand(topCmd)
}
