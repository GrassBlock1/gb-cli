package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// portableCmd represents the portable command
var portableCmd = &cobra.Command{
	Use:   "portable",
	Short: "Make any program run like a portable one (Linux only)",
	Long:  `Make any program run like a portable one. It sets the HOME variable to ./Data, and link some configs there`,
	Args:  cobra.MinimumNArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return nil, cobra.ShellCompDirectiveFilterFileExt
	},
	Run: func(cmd *cobra.Command, args []string) {
		portable := exec.Command(args[0], args[1:]...)
		portable.Env = os.Environ()
		err := os.MkdirAll("./Data", 0755)
		if err != nil {
			log.Fatal("Error while creating isolated $HOME", err)
		}
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal("Error while getting $PWD", err)
		}
		portable.Env = append(portable.Env, "HOME="+pwd+"/Data")
		cfgDir, err := os.UserConfigDir()
		if err != nil {
			log.Fatal("Error while getting config dir", err)
		}
		// TODO: restrict to specific config
		err = os.Symlink(cfgDir, pwd+"/Data/.config")
		if err != nil {
			log.Fatal("Error while linking config dir", err)
		}
		portable.Stdout = os.Stdout
		portable.Stderr = os.Stderr
		portable.Stdin = os.Stdin
		err = portable.Run()
		if err != nil {
			log.Fatal("Error while running command: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(portableCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// portableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// portableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
