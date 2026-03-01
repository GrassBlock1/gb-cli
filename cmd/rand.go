package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// randCmd represents the rand command
var randCmd = &cobra.Command{
	Use:   "rand",
	Short: "Get real random number from multiple sources",
	Long:  `Get real random number from random.org (default) / generate from drand / nist / qrng beacon`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		min, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: ")
			os.Exit(1)
		}
		max, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: ")
			os.Exit(1)
		}
		if min >= max {
			os.Exit(1)
		}
		num, err := getFromRandomOrg(min, max)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(num)
	},
}

func init() {
	rootCmd.AddCommand(randCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getFromRandomOrg(min, max int) (int, error) {
	url := fmt.Sprintf("https://www.random.org/integers/?num=1&min=%d&max=%d&col=1&base=10&format=plain&rnd=new", min, max)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("server not ok")
	}
	var num int
	_, err = fmt.Fscanf(resp.Body, "%d", &num)
	if err != nil {
		return 0, err
	}
	return num, nil
}

//func getRandomSeed(from string) (string, error) {
//
//}
