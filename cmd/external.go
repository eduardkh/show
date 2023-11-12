package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// externalCmd represents the external command
var externalCmd = &cobra.Command{
	Use:   "external",
	Short: "Get your external IP Address",
	Long: `
	Get your external IP Address`,
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://ifconfig.me/ip"
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Accept", "application/json")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		raw, err := cmd.Flags().GetBool("raw")
		if err != nil {
			fmt.Println(err)
			return
		}

		if raw {
			fmt.Println(string(body))
		} else {
			fmt.Println("Your external IP is:", string(body))
		}
	},
}

func init() {
	ipCmd.AddCommand(externalCmd)
	externalCmd.Flags().BoolP("raw", "r", false, "get IP address only")
}
