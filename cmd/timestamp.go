package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// timestampCmd represents the timestamp command
var timestampCmd = &cobra.Command{
	Use:   "timestamp",
	Short: "Get a timestamp",
	Long:  `Get a timestamp`,
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		flg, _ := cmd.Flags().GetBool("epoch")
		if flg {
			fmt.Println(now.Unix())
		} else {
			fmt.Println(now)
		}
	},
}

func init() {
	rootCmd.AddCommand(timestampCmd)
	timestampCmd.Flags().BoolP("epoch", "e", false, "get epoch timestamp")
}
