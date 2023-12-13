package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// IPInfo represents the structure of the response from IPinfo.io
type IPInfo struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
}

// fetchIPInfo fetches information about an IP address
func fetchIPInfo(ip string, orgOnly bool) {
	resp, err := http.Get(fmt.Sprintf("https://ipinfo.io/%s", ip))
	if err != nil {
		fmt.Println("Error fetching IP information:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var ipInfo IPInfo
	err = json.Unmarshal(body, &ipInfo)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	if orgOnly {
		fmt.Println("Organization: " + ipInfo.Org)
	} else {
		fmt.Printf("IP Address  : %+v\n", ipInfo.IP)
		fmt.Printf("Hostname    : %+v\n", ipInfo.Hostname)
		fmt.Printf("City        : %+v\n", ipInfo.City)
		fmt.Printf("Region      : %+v\n", ipInfo.Region)
		fmt.Printf("Country     : %+v\n", ipInfo.Country)
		fmt.Printf("Location    : %+v\n", ipInfo.Loc)
		fmt.Printf("Organization: %+v\n", ipInfo.Org)
	}
}

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info [IP address]",
	Short: "Get information about an IP address",
	Long:  `Get detailed information about an IP address using IPinfo.io.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Usage()
			return
		}
		ip := args[0]
		orgOnly, _ := cmd.Flags().GetBool("organization")
		fetchIPInfo(ip, orgOnly)
	},
}

func init() {
	ipCmd.AddCommand(infoCmd)
	infoCmd.Flags().BoolP("organization", "o", false, "Display only the organization of the IP address")
}
