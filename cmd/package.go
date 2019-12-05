package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

// packageCmd represents the package command
var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "Packages your current branch or one passed with the -b flag for deployment.",
	Long: `Runs the process of packaging a release for any repo using git flow standards. This will do the following:
* Take either the current branch or the branch name passed with the -b flag and merge it into master if CI passed.
* Tag the commit where the aforementioned branch was merged into master.
* Push the tag and branch to origin
* Post a formatted message to the #releases slack channel
For example: sb-release -b Release/19.69.0`,
	Run: func(cmd *cobra.Command, args []string) {
		branch, _ := cmd.Flags().GetString("branch")

		if branch == "" {
			branch = "currentBranch"
		}
		fmt.Println("Prepare release for " + branch + "?")
		gitBranchName()
	},
}

func init() {
	rootCmd.AddCommand(packageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// packageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	packageCmd.Flags().BoolP("branch", "b", false, "Branch name to prepare a release for.")
}

func gitBranchName() {
	command := exec.Command("git rev-parse --abbrev-ref HEAD")
	err := command.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
