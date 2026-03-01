package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get things from the Internet",
	Long:  `Get quotes/hitokoto/dev excuses`,
}

var hitokotoCmd = &cobra.Command{
	Use:   "hitokoto",
	Short: "Get sentence from hitokoto.cn",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getSentenceFromHitokoto())
	},
}

var selfQuotesCmd = &cobra.Command{
	Use:   "quotes",
	Short: "Get sentence from self collection",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getSentenceFromSelf())
	},
}

var devExcusesCmd = &cobra.Command{
	Use:   "dev-excuses",
	Short: "Get sentence from developer excuses (api from tabliss)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getSentenceFromDevExcuses())
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(hitokotoCmd)
	getCmd.AddCommand(selfQuotesCmd)
	getCmd.AddCommand(devExcusesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//getCmd.Flags().BoolP("cache", "c", false, "Also cache quotes to disk")
}

func getSentenceFromHitokoto() string {
	resp, err := http.Get("https://v1.hitokoto.cn")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var data interface{}
	e := json.NewDecoder(resp.Body).Decode(&data)
	if e != nil {
		panic(e)
	}
	dMap := data.(map[string]interface{})
	sentence := dMap["hitokoto"].(string)
	from := dMap["from"].(string)
	// format sentence a little
	result := fmt.Sprintf("%s\n %s\t————%s", sentence, strings.Repeat(" ", utf8.RuneCountInString(sentence)*2), from) // our sentence is always utf-8, so we need to calculate the characters differently
	return result
}

func getSentenceFromSelf() string {
	resp, err := http.Get("https://gb0.dev/g/quotes")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var data interface{}
	e := json.NewDecoder(resp.Body).Decode(&data)
	if e != nil {
		panic(e)
	}
	dMap := data.(map[string]interface{})
	sentence := dMap["sentence"].(string)
	from := ""
	// somewhat fallback algorithm
	if dMap["cite"] != nil {
		from = dMap["cite"].(string)
	} else if dMap["author"] != nil {
		from = dMap["author"].(string)
	} else if dMap["source"] != nil {
		from = dMap["source"].(string)
	} else {
		from = "佚名"
	}
	// format sentence a little
	result := fmt.Sprintf("%s\n %s\t————%s", sentence, strings.Repeat(" ", utf8.RuneCountInString(sentence)*2), from)
	return result
}

func getSentenceFromDevExcuses() string {
	resp, err := http.Get("https://api.tabliss.io/v1/developer-excuses")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var data interface{}
	e := json.NewDecoder(resp.Body).Decode(&data)
	if e != nil {
		panic(e)
	}
	dMap := data.(map[string]interface{})
	sentence := dMap["data"].(string)
	// format sentence a little
	result := fmt.Sprintf("%s\n %s\t————%s", sentence, strings.Repeat(" ", len(sentence)), "Developer excuses")
	return result
}
