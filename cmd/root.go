package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "perro",
	Short: "USAGE: mergectl [source branch] [target branch]",
	Long:  `mergectl is a tool of "git merge". This command merges multiple git branches.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello perro!")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
