package cmd

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// rngCmd represents the rng command
var rngCmd = &cobra.Command{
	Use:   "rng",
	Short: "generate random number like fish",
	Long:  `rng generates a cryptographically secure pseudo-random integer from a uniform distribution. It can be used like fish's internal function without seeding.`,
	Run: func(cmd *cobra.Command, args []string) {
		var Reader io.Reader = rand.Reader
		switch len(args) {
		case 0:
			n, err := rand.Int(Reader, big.NewInt(32768))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(n)
		case 1:
			if args[1] == "choice" {
				log.Fatal("Error: no items to be selected")
			}
			fmt.Println("Seeding is not allowed here")
			os.Exit(1)
		case 2:
			minN, errMin := strconv.Atoi(args[0])
			if errMin != nil {
				log.Fatal("Error: invalid minium value")
			}
			maxN, errMax := strconv.Atoi(args[1])
			if errMax != nil {
				log.Fatal("Error: invalid maximum value")
			}
			if minN >= maxN {
				log.Fatal("Error: max value must be bigger than min value")
			}
			rangeSize := maxN - minN + 1
			result, err := rand.Int(Reader, big.NewInt(int64(rangeSize)))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(result.Int64() + int64(minN))
		case 3:
			if args[0] == "choice" {
				n, err := rand.Int(Reader, big.NewInt(2))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(args[n.Int64()+1])
			}
			start, errMin := strconv.Atoi(args[0])
			if errMin != nil {
				log.Fatal("Error: invalid minium value")
			}
			step, errStep := strconv.Atoi(args[1])
			if errStep != nil {
				log.Fatal("Error: invalid step value")
			}
			end, errMax := strconv.Atoi(args[2])
			if errMax != nil {
				log.Fatal("Error: invalid maximum value")
			}
			if step <= 0 {
				log.Fatal("Error: step must be positive")
			}
			if start >= end {
				log.Fatal("Error: max value must be bigger than min value")
			}
			count := (end-start)/step + 1
			index, err := rand.Int(Reader, big.NewInt(int64(count)))
			if err != nil {
				log.Fatal(err)
			}
			result := start + int(index.Int64())*step
			fmt.Println(result)
		default:
			if args[0] == "choice" && len(args) > 1 {
				n, err := rand.Int(Reader, big.NewInt(int64(len(args)-1)))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(args[n.Int64()+1])
			} else {
				log.Fatal("Error: Too many arguments. see random in fish shell for usage.")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rngCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rngCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rngCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
