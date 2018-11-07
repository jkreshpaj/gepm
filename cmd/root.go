package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gepm",
	Short: "run gepm [arg] to search for packages or gepm to install packages from file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: Unable to find packages.json on current directory")
		}

		if len(args) > 0 {
			done := make(chan bool)
			s.Prefix = "Searching for " + args[0] + " "
			s.Start()
			go searchPackages(args[0], done)
			<-done
			s.Stop()
			if len(packages.Hits) > 0 {
				makePrompt()
			}
			if len(packages.Hits) == 0 {
				fmt.Println("No packages named", args[0], "found")
			}
		}
	},
}

//Execute gemp
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
