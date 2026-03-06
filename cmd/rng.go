package cmd

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"math/big"
	"strconv"

	"github.com/spf13/cobra"
)

// rngCmd represents the rng command
var rngCmd = &cobra.Command{
	Use:   "rng",
	Short: "generate random number like fish",
	Long:  `rng generates a cryptographically secure pseudo-random integer from a uniform distribution. It can be used like fish's internal function without seeding.`,
	Args: func(cmd *cobra.Command, args []string) error {
		switch len(args) {
		case 0:
			return nil
		case 1:
			if args[0] == "choice" {
				return fmt.Errorf("nothing to choose from")
			}
			return fmt.Errorf("seeding is not allowed here")
		case 2:
			if args[0] == "choice" {
				return nil
			}
			_, errMin := strconv.Atoi(args[0])
			if errMin != nil {
				return fmt.Errorf("invalid minium value")
			}
			_, errMax := strconv.Atoi(args[1])
			if errMax != nil {
				return fmt.Errorf("invalid maximum value")
			}
			return nil
		case 3:
			if args[0] == "choice" {
				return nil
			}
			_, errMin := strconv.Atoi(args[0])
			if errMin != nil {
				return fmt.Errorf("invalid minium value")
			}
			_, errStep := strconv.Atoi(args[1])
			if errStep != nil {
				return fmt.Errorf("invalid step value")
			}
			_, errMax := strconv.Atoi(args[2])
			if errMax != nil {
				return fmt.Errorf("invalid maximum value")
			}
			return nil
		default:
			if args[0] == "choice" {
				return nil
			}
			return fmt.Errorf("too many arguments. see random in fish shell for usage")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var Reader io.Reader = rand.Reader
		switch len(args) {
		case 0:
			n, err := rand.Int(Reader, big.NewInt(32768))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(n)
		case 2:
			if args[0] == "choice" {
				fmt.Println(args[1])
				return
			}
			minN, _ := strconv.Atoi(args[0])
			maxN, _ := strconv.Atoi(args[1])
			if minN >= maxN {
				log.Fatal("max value must be bigger than min value")
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
			start, _ := strconv.Atoi(args[0])
			step, _ := strconv.Atoi(args[1])
			end, _ := strconv.Atoi(args[2])
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
