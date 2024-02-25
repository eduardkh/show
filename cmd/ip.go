package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "List all IP commands",
	Long:  `List all IP commands`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Usage(); err != nil {
			// Handle the error, e.g., log it or output to stderr
			fmt.Fprintf(os.Stderr, "Error displaying usage: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)
}
