package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/spf13/cobra"
)

// matrixCmd represents the matrix command
var matrixCmd = &cobra.Command{
	Use:   "matrix <domain>",
	Short: "Get information of a matrix homeserver",
	Long:  `Get a server domain behind a matrix homeserver and the software powers it`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("too many arguments")
		} else if !isValidDomain(args[0]) {
			return fmt.Errorf("invalid domain format")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		domain := args[0]
		homeserver, software, version := getServerInfo(domain)
		fmt.Printf("%s is behind %s, running %s (%s).", domain, homeserver, software, version)
	},
}

func init() {
	rootCmd.AddCommand(matrixCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// matrixCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// matrixCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func isValidDomain(domain string) bool {
	// Pattern matches basic domain format: name.extension or subdomain.name.extension
	pattern := `^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, domain)
	return matched
}

func getServerInfo(domain string) (string, string, string) {
	wellKnownURL := fmt.Sprintf("https://%s/.well-known/matrix/client", domain)
	resp, err := http.Get(wellKnownURL)
	if err != nil {
		log.Fatal("Error requesting well-known URL: ", err)
	}
	defer resp.Body.Close()
	var data interface{}
	e := json.NewDecoder(resp.Body).Decode(&data)
	if e != nil {
		log.Fatal("Error while parsing well-known JSON: ", e)
	}
	dMap := data.(map[string]interface{})
	homeserverData := dMap["m.homeserver"].(map[string]interface{})
	homeserver := homeserverData["base_url"].(string)
	// get from federation api
	federationURL := fmt.Sprintf("%s/_matrix/federation/v1/version", homeserver)
	fresp, ferr := http.Get(federationURL)
	if ferr != nil {
		log.Fatal(err)
	}
	defer fresp.Body.Close()
	var softwareData interface{}
	fe := json.NewDecoder(fresp.Body).Decode(&softwareData)
	if fe != nil {
		log.Fatal(fe)
	}
	softwareDataMap := softwareData.(map[string]interface{})
	// dealing with nesting data
	serverMap := softwareDataMap["server"].(map[string]interface{})
	software := serverMap["name"].(string)
	version := serverMap["version"].(string)
	return homeserver, software, version
}
