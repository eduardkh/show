package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "show",
	Short: "The show command displays the current IP and networking configuration of the system.",
	Long: `Use the show command to view the system's IP and networking details. Subcommands include:

show ip: Displays IP configuration.
show dns: Lists configured DNS servers.
show interfaces: Shows network interfaces and their status.
This command is read-only and safe for troubleshooting network configurations.`,
}

// Execute runs the root command.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
