package cmd

import (
	"encoding/base32"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// b32Cmd represents the b32 command
var b32Cmd = &cobra.Command{
	Use:   "b32",
	Short: "Encode/decode base32 string",
	Long:  `A basic base64 decoder/encoder, which takes a string as a argument`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := strings.Join(args, " ")
		decode, _ := cmd.Flags().GetBool("decode")
		if decode {
			decodedString, err := base32.StdEncoding.DecodeString(input)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(decodedString))
		} else {
			fmt.Println(base32.StdEncoding.EncodeToString([]byte(input)))
		}
	},
}

func init() {
	rootCmd.AddCommand(b32Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// b32Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	b32Cmd.Flags().BoolP("decode", "d", false, "decode the provided string")
}
