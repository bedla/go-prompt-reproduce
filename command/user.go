package command

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:     "user [no options!]",
	Short:   "user setup",
	Aliases: []string{"u",},
	Run: func(cmd *cobra.Command, args []string) {
		setupUsername()
	},
}

func setupUsername() {
	usernamePrompt := promptui.Prompt{
		Label: "User name",
	}
	username, err := usernamePrompt.Run()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("user is " + username)
	}
}

func init() {
	rootCmd.AddCommand(userCmd)
}
