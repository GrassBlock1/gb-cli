package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var verbosity int
var tumbleweed = `
      .   '     .
    '   .  _  '   .
   .  '  /   \  '  .
   '  .  \ _ /  .  '
      .    '   .
`

// repulseCmd represents the repulse command
var repulseCmd = &cobra.Command{
	Use:                "repulse",
	Short:              "Get super repulsive powers",
	DisableSuggestions: true,
	Hidden:             true,
	Aliases:            []string{"repel"},
	Run: func(cmd *cobra.Command, args []string) {
		// imitate from aptitude
		// Thanks to Gemini for tumbleweed idea... Although I haven't used OpenSUSE tumbleweed before.
		// And for its review. I have a bad English...
		switch verbosity {
		case 0:
			fmt.Println("There are no Easter eggs in this program. Just another lonely project with no users.")
		case 1:
			fmt.Println("There really are no Easter eggs in this program. They're as invisible as the bass guitar in a song.")
		case 2:
			fmt.Println("I thought I told you there are no Easter eggs here.")
		case 3:
			fmt.Println("Seriously, could you please stop trying?")
		case 4:
			fmt.Println("Okay... if I give you an Easter egg, will you go away?")
		case 5:
			fmt.Printf("Fine... Here you go:\n %s\n", tumbleweed)
		default:
			fmt.Println("Whether you ask for it or not, it's still just a tumbleweed.")
		}
	},
}

func init() {
	rootCmd.AddCommand(repulseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repulseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repulseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	repulseCmd.Flags().CountVarP(&verbosity, "verbose", "v", "verbose output")
}
