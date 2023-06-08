package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "show",
	Short: "The show command displays the current IP and networking configuration of the system.",
	Long: `Use the show command to view the system's IP and networking details. Subcommands include:

show ip: Displays IP configuration.
show dns: Lists configured DNS servers.
show interfaces: Shows network interfaces and their status.
This command is read-only and safe for troubleshooting network configurations.`,
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
