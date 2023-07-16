package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hnCLI",
	Short: "A simple CLI to interact with HN",
	Long:  "Lorem Ipsum Dolor",

	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, "Encountered error:", err)
		os.Exit(1)
	}
}
