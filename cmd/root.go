package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gb",
	Short:   "gb is a cli tool for everyday tasks",
	Long:    "gb is a cli tool for everyday tasks - base32/64/85 encoding/decoding, getting quotes, etc",
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			return
		}
		fmt.Println("\t\t\t\tThis cli has Super Repulsive Powers.")
	},
}

func Execute() {
	rootCmd.Execute()
}
