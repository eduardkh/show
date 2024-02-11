package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// AbuseIPDBResponse represents the JSON response structure from AbuseIPDB API
type AbuseIPDBResponse struct {
	Data struct {
		IPAddress            string   `json:"ipAddress"`
		IsPublic             bool     `json:"isPublic"`
		IpVersion            int      `json:"ipVersion"`
		IsWhitelisted        *bool    `json:"isWhitelisted"` // Pointer to handle null
		AbuseConfidenceScore int      `json:"abuseConfidenceScore"`
		CountryCode          string   `json:"countryCode"`
		UsageType            string   `json:"usageType"`
		Isp                  string   `json:"isp"`
		Domain               string   `json:"domain"`
		Hostnames            []string `json:"hostnames"`
		IsTor                bool     `json:"isTor"`
		CountryName          string   `json:"countryName"`
		TotalReports         int      `json:"totalReports"`
		NumDistinctUsers     int      `json:"numDistinctUsers"`
		LastReportedAt       *string  `json:"lastReportedAt"` // Pointer to handle null
		Reports              []struct {
			ReportedAt          string `json:"reportedAt"`
			Comment             string `json:"comment"`
			Categories          []int  `json:"categories"`
			ReporterId          int    `json:"reporterId"`
			ReporterCountryCode string `json:"reporterCountryCode"`
			ReporterCountryName string `json:"reporterCountryName"`
		} `json:"reports"`
	} `json:"data"`
}

var (
	configPath = filepath.Join(os.Getenv("APPDATA"), "show", "config.yaml")
)

func ensureConfigFile() error {
	// Ensure the config directory exists
	if err := os.MkdirAll(filepath.Dir(configPath), os.ModePerm); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Initialize viper with the config file
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	// Check if the config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// The config file does not exist, attempt to create it
		file, err := os.Create(configPath)
		if err != nil {
			return fmt.Errorf("failed to directly create config file: %w", err)
		}
		file.Close() // Close the file after creating it

		// Set a default value for "api_key" and attempt to write to the file
		viper.Set("api_key", "")
		if err := viper.WriteConfig(); err != nil {
			return fmt.Errorf("failed to write default config: %w", err)
		}
	} else if err != nil {
		// An error occurred trying to check the file, unrelated to file not existing
		return fmt.Errorf("error checking config file: %w", err)
	} else {
		// The config file exists, try reading it
		if err := viper.ReadInConfig(); err != nil {
			return fmt.Errorf("failed to read config file: %w", err)
		}
	}

	return nil
}

func getAPIKey() (string, error) {
	if err := ensureConfigFile(); err != nil {
		return "", err
	}
	return viper.GetString("api_key"), nil
}

func updateAPIKey(newKey string) error {
	if err := ensureConfigFile(); err != nil {
		return err
	}
	viper.Set("api_key", newKey)
	return viper.WriteConfig()
}

func queryAbuseIPDB(ipAddress, apiKey string) (*AbuseIPDBResponse, error) {
	url := fmt.Sprintf("https://api.abuseipdb.com/api/v2/check?maxAgeInDays=90&verbose&ipAddress=%s", ipAddress)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Key", apiKey)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result AbuseIPDBResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func printAbuseIPDBResponse(response *AbuseIPDBResponse) {
	fmt.Printf("IP Address            : %+v\n", response.Data.IPAddress)
	fmt.Printf("Abuse Confidence Score: %+v\n", response.Data.AbuseConfidenceScore)
	fmt.Printf("Country Code          : %+v\n", response.Data.CountryCode)
	fmt.Printf("Usage Type            : %+v\n", response.Data.UsageType)
	if len(response.Data.Hostnames) > 0 {
		fmt.Printf("Hostnames             : %+v\n", response.Data.Hostnames)
	}
	if response.Data.Domain != "" {
		fmt.Printf("Domain                : %+v\n", response.Data.Domain)
	}
	if response.Data.Isp != "" {
		fmt.Printf("ISP                   : %+v\n", response.Data.Isp)
	}
	fmt.Printf("Is TOR Server         : %+v\n", response.Data.IsTor)
	if response.Data.LastReportedAt != nil {
		fmt.Printf("Last Reported At      : %+v\n", *response.Data.LastReportedAt)
	}

	// For reports, it's best to loop over individual responses.
	if len(response.Data.Reports) > 0 {
		fmt.Println("")
		fmt.Println("Reports:")
		for _, report := range response.Data.Reports {
			fmt.Println("\n🕙 Reported At     : ", report.ReportedAt)
			fmt.Println("🌍 Reporter Country: ", report.ReporterCountryName)
			fmt.Println("📝 Comment         : ", report.Comment)
		}
	} else {
		fmt.Println("No reports available.")
	}
}

// abuseipdbCmd represents the abuseipdb command
var abuseipdbCmd = &cobra.Command{
	Use:   "abuseipdb [IP address]",
	Short: "Query abuse information for an IP address from AbuseIPDB",
	Long: `Query abuse information for an IP address using the AbuseIPDB API.

Example usage:
show ip abuseipdb 8.8.8.8`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if the update-key flag is set
		updateKey, _ := cmd.Flags().GetString("update-key")
		if updateKey != "" {
			if err := updateAPIKey(updateKey); err != nil {
				fmt.Println("Error updating API key:", err)
				return
			}
			fmt.Println("API key updated successfully.")
			return
		}
		// Check if arguments are provided
		if len(args) == 0 {
			cmd.Usage()
			fmt.Println("\nUsage:")
			fmt.Println("show ip abuseipdb 8.8.8.8")
			return
		}

		ipAddress := args[0] // Take the first argument as the IP address

		apiKey, err := getAPIKey()
		if err != nil {
			fmt.Println("Error reading API key:", err)
			return
		}

		response, err := queryAbuseIPDB(ipAddress, apiKey)
		if err != nil {
			fmt.Println("Error querying AbuseIPDB:", err)
			return
		}

		printAbuseIPDBResponse(response)

	},
}

func init() {
	abuseipdbCmd.Flags().StringP("update-key", "u", "", "Update the API key in the configuration file")
	ipCmd.AddCommand(abuseipdbCmd)
}
