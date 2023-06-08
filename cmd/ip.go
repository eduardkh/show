package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "List all IP commands",
	Long:  `List all IP commands`,
	Run: func(cmd *cobra.Command, args []string) {
		commands := []string{
			"Usage:",
			"show ip external",
			"show ip interface",
			"show ip interface brief",
		}
		for _, command := range commands {
			fmt.Println(command)
		}
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)
}
