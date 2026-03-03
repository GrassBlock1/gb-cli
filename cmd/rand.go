package cmd

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"slices"
	"strconv"

	"github.com/spf13/cobra"
)

// randCmd represents the rand command
var randCmd = &cobra.Command{
	Use:   "rand",
	Short: "Get 'true', verifiable random number from multiple sources",
	Long:  `Get 'true', verifiable random number from random.org (default) / generate from drand / nist / qrng beacon`,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
		if len(args) == 0 {
			completion := []cobra.Completion{
				cobra.CompletionWithDesc("drand", "use seed from drand.love"),
				cobra.CompletionWithDesc("nist", "use seed from csrc.nist.gov beacon"),
				cobra.CompletionWithDesc("qrng", "use seed from anu qrng")}
			return completion, cobra.ShellCompDirectiveNoFileComp
		}
		return nil, cobra.ShellCompDirectiveNoFileComp
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 2 {
			_, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid min value")
			}
			_, errMax := strconv.Atoi(args[1])
			if errMax != nil {
				return fmt.Errorf("invalid max value")
			}
			return nil
		} else if len(args) == 3 {
			methods := []string{"drand", "nist", "qrng"}
			if !slices.Contains(methods, args[0]) {
				return fmt.Errorf("invalid method")
			}
			_, err := strconv.Atoi(args[1])
			if err != nil {
				return fmt.Errorf("invalid min value")
			}
			_, errMax := strconv.Atoi(args[2])
			if errMax != nil {
				return fmt.Errorf("invalid max value")
			}
			return nil
		}
		return fmt.Errorf("unexpected arguments")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			minN, _ := strconv.Atoi(args[0])
			maxN, _ := strconv.Atoi(args[1])
			if minN >= maxN {
				log.Fatal("Error: min value must be bigger than max value")
			}
			num, err := getFromRandomOrg(minN, maxN)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			fmt.Println(num)
		} else if len(args) == 3 {
			source := args[0]
			seed, err := getRandomSeed(source)
			if err != nil {
				log.Fatal("error while getting seed:", err)
			}
			minN, _ := strconv.Atoi(args[1])
			maxN, _ := strconv.Atoi(args[2])
			if minN >= maxN {
				log.Fatal("Error: min value must be bigger than max value")
			}
			// TODO: also display the seed and other info
			fmt.Println(getFromSeed(seed, minN, maxN))
		}
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
		log.Fatal(err)
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

func getRandomSeed(from string) (string, error) {
	if from == "drand" {
		resp, err := http.Get("https://api.drand.sh/public/latest")
		if err != nil {
			log.Fatal("Error getting random seed from drand", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return "", fmt.Errorf("server not ok")
		}
		var data interface{}
		e := json.NewDecoder(resp.Body).Decode(&data)
		if e != nil {
			log.Fatal("Error while parsing well-known JSON: ", e)
		}
		dataMap := data.(map[string]interface{})
		return dataMap["randomness"].(string), nil
	}
	if from == "nist" {
		resp, err := http.Get("https://beacon.nist.gov/beacon/2.0/pulse/last")
		if err != nil {
			log.Fatal("Error getting random seed from nist", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return "", fmt.Errorf("server not ok")
		}
		var data interface{}
		e := json.NewDecoder(resp.Body).Decode(&data)
		if e != nil {
			log.Fatal("Error while parsing well-known JSON: ", e)
		}
		dataMap := data.(map[string]interface{})
		pulseMap := dataMap["pulse"].(map[string]interface{})
		return pulseMap["outputValue"].(string), nil
	}
	if from == "qrng" {
		resp, err := http.Get("https://qrng.anu.edu.au/wp-content/plugins/colours-plugin/get_block_alpha.php")
		if err != nil {
			log.Fatal("Error getting random seed from nist", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return "", fmt.Errorf("server not ok")
		}
		var data string
		_, e := fmt.Fscanf(resp.Body, "%s", &data)
		if e != nil {
			return "", err
		}
		return data, nil
	}
	return "", fmt.Errorf("error: unknown source")
}

func getFromSeed(seed string, min, max int) int {
	seedBytes, _ := hex.DecodeString(seed)
	h := sha256.Sum256(seedBytes)

	// init as a seed
	trueRand := rand.New(rand.NewChaCha8(h))
	return trueRand.IntN(max-min+1) + min
}
