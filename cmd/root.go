package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	search   = spinner.New(spinner.CharSets[7], 100*time.Millisecond)
	progress = spinner.New(spinner.CharSets[7], 100*time.Millisecond)
	success  = color.New(color.FgGreen)
	failed   = color.New(color.FgRed)
	packages = new(Packages)
)

var rootCmd = &cobra.Command{
	Use:   "gepm",
	Short: "run gepm [arg] to search for packages or gepm to install packages from file",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			failed.Println("Error: Unable to find packages.json on current directory")
		}

		if len(args) > 0 {
			done := make(chan bool)
			search.Prefix = "Searching for " + args[0] + " "

			search.Start()
			go searchPackages(args[0], done)
			<-done
			search.Stop()

			if len(packages.Hits) > 0 {
				makePrompt()
			}
			if len(packages.Hits) == 0 {
				failed.Println("No packages named", args[0], "found")
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
