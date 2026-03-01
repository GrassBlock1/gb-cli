package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gb",
	Short: "gb is a cli tool for everyday tasks",
	Long:  "gb is a cli tool for everyday tasks - base32/64/85 encoding/decoding, getting quotes, etc",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			return
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing '%s'\n", err)
		os.Exit(1)
	}
}
