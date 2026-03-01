package cmd

import (
	"crypto/rand"
	_ "embed"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	//go:embed assets/eff-wordlist.txt
	wordlist string
	symbols  = "1234567890`-=[];,./|\\~!@#$%^&*()_+{}:<>?\""
)

// passwdCmd represents the passwd command
var passwdCmd = &cobra.Command{
	Use:   "passwd",
	Short: "Generate a memorable secure password",
	Long:  `Generate a memorable secure password using EFF's word list and random numbers`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		chars, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid input")
			os.Exit(1)
		}
		if chars < 8 {
			fmt.Fprintln(os.Stderr, "Error: password must be at least 8 characters")
			os.Exit(1)
		}
		fmt.Println(genPasswd(chars))
	},
}

func init() {
	rootCmd.AddCommand(passwdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// passwdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// passwdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getRandomInt(max int) int {
	if max <= 0 {
		return 0
	}
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return int(n.Int64())
}

func genPasswd(chars int) string {
	var passwd strings.Builder
	wordlistData := strings.Split(strings.TrimSpace(wordlist), "\n")
	firstWord := wordlistData[getRandomInt(len(wordlistData))]
	if getRandomInt(2) == 1 {
		firstWord = cases.Title(language.English, cases.Compact).String(firstWord)
	}
	passwd.WriteString(firstWord)
	for {
		remaining := chars - passwd.Len()
		if remaining < 7 {
			break
		}
		// get random fillings first
		filling := getRandomInt(3) + 1
		for i := 0; i < filling; i++ {
			if passwd.Len() < chars {
				passwd.WriteByte(symbols[getRandomInt(len(symbols))])
			}
		}
		// move to the next one
		nextWord := wordlistData[getRandomInt(len(wordlistData))]
		if getRandomInt(2) == 1 {
			nextWord = cases.Title(language.English, cases.Compact).String(nextWord)
		}
		if passwd.Len()+len(nextWord) > chars {
			break
		}
		passwd.WriteString(nextWord)
	}
	for passwd.Len() < chars {
		passwd.WriteByte(symbols[getRandomInt(len(symbols))])
	}
	return passwd.String()
}
