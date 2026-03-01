package cmd

import (
	"encoding/ascii85"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// b85Cmd represents the b85 command
var b85Cmd = &cobra.Command{
	Use:   "b85",
	Short: "Encode/decode standard base85(ascii85) string",
	Long:  `A basic base85(ascii85) decoder/encoder, which takes a string as a argument`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := strings.Join(args, " ")
		decode, _ := cmd.Flags().GetBool("decode")
		if decode {
			dst := make([]byte, len(input)*2)
			n, _, err := ascii85.Decode(dst, []byte(input), true)
			if err != nil {
				return
			}
			fmt.Println(string(dst[:n]))
		} else {
			// in case the index is out of range
			dst := make([]byte, len(input)*2)
			n := ascii85.Encode(dst, []byte(input))
			fmt.Println(string(dst[:n]))
		}
	},
}

func init() {
	rootCmd.AddCommand(b85Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// b85Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	b85Cmd.Flags().BoolP("decode", "d", false, "decode the provided string")
}
