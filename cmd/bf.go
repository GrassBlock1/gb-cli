package cmd

import (
	"git.gay/gb/gb/internal/bf"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// bfCmd represents the bf command
var bfCmd = &cobra.Command{
	Use:   "bf",
	Short: "run a brainfuck script",
	Long:  `Run a literal brainfuck script, or eval one from file`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath, _ := cmd.Flags().GetString("file")
		var scriptReader *strings.Reader

		if filePath != "" {
			content, err := os.ReadFile(filePath)
			if err != nil {
				log.Fatal(err)
			}
			scriptReader = strings.NewReader(string(content))
		} else if len(args) == 0 {
			log.Fatal("Error: no script or file path provided")
		}
		script := args[0]
		scriptReader = strings.NewReader(script)
		err := bf.Eval(scriptReader, os.Stdin, os.Stderr)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(bfCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bfCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bfCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	bfCmd.Flags().StringP("file", "f", "", "Running script from a file")
}
