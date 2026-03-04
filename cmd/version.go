package cmd

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"runtime"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var commit = func() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return "g" + setting.Value[0:8]
			}
		}
	}

	return "devel"
}()

var Version = "0.0.0-" + commit

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version and more info",
	Long:  `Prints the version and more info`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gb version %s (on %s %s)\n", Version, runtime.GOOS, runtime.GOARCH)
		reader := rand.Reader
		n, _ := rand.Int(reader, big.NewInt(1000))
		if n.Int64() == 0 {
			fmt.Printf("\nCake is a lie.\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
