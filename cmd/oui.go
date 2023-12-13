/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// DownloadOUIFile downloads the OUI CSV file from the IEEE website and stores it in the local application data directory.
func DownloadOUIFile() error {
	// Determine the local application data directory.
	appDataDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("error getting user config directory: %w", err)
	}

	// Create a subdirectory for your app.
	appDataDir = filepath.Join(appDataDir, "show")
	if err := os.MkdirAll(appDataDir, 0755); err != nil {
		return fmt.Errorf("error creating application data directory: %w", err)
	}

	// Define the URL and the local file path.
	ouiURL := "http://standards-oui.ieee.org/oui/oui.csv"
	localFilePath := filepath.Join(appDataDir, "oui.csv")

	// Download the OUI file.
	err = downloadFile(localFilePath, ouiURL)
	if err != nil {
		return fmt.Errorf("error downloading the OUI file: %w", err)
	}

	fmt.Printf("OUI file downloaded successfully to %s\n", localFilePath)
	return nil
}

// downloadFile downloads the file from the given URL to the given local path.
func downloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// PrintOUIInfo searches the OUI CSV file for the given MAC address and prints its information.
func PrintOUIInfo(macAddress string) error {
	// Standardize the MAC address by removing delimiters and converting to uppercase
	macAddress = strings.ToUpper(macAddress)
	macAddress = strings.NewReplacer(":", "", "-", "", ".", "").Replace(macAddress)

	// Now that we've cleaned the MAC address, check if it's at least 6 characters long
	if len(macAddress) < 6 {
		return fmt.Errorf("MAC address must be at least 6 characters after standardizing")
	}

	// Trim the MAC address to the first 6 characters (OUI portion)
	ouiPortion := macAddress[:6]

	// Get the local application data directory
	appDataDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	// Construct the path to the oui.csv file
	csvFilePath := filepath.Join(appDataDir, "show", "oui.csv")

	// Open the CSV file
	file, err := os.Open(csvFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read and ignore the header line
	if _, err := reader.Read(); err != nil {
		return err
	}

	// Iterate through the CSV records and print the matching OUI information
	found := false
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		// Check if the current record contains the OUI portion of the MAC address
		if strings.HasPrefix(strings.ToUpper(record[1]), ouiPortion) {
			fmt.Printf("MAC Address: \"%s\"\n", macAddress)
			fmt.Printf("Vendor: \"%s\"\n", record[2])
			found = true
			break // Stop after finding the first match
		}
	}

	if !found {
		fmt.Println("No matching records found for the given MAC address.")
	}

	return nil
}

// ouiCmd represents the oui command
var ouiCmd = &cobra.Command{
	Use:   "oui",
	Short: "Display vendor information for a given MAC address based on its OUI.",
	Long: `
Display vendor information for a given MAC address based on its OUI.

Usage:
show ip oui 005056C00001
show ip oui 00:50:56:C0:00:01
show ip oui 00-50-56-C0-00-01
show ip oui '0050.56C0.0001'
	`,
	Run: func(cmd *cobra.Command, args []string) {
		update, _ := cmd.Flags().GetBool("update")
		if update {
			// Existing logic for updating the OUI file

			err := DownloadOUIFile()
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			// Check if a MAC address argument is provided
			if len(args) == 0 {
				cmd.Usage()
				fmt.Println(`
Usage:
show ip oui 005056C00001
show ip oui 00:50:56:C0:00:01
show ip oui 00-50-56-C0-00-01
show ip oui '0050.56C0.0001'
				`)
				return
			}
			macAddress := args[0] // assuming the MAC address is the first argument
			err := PrintOUIInfo(macAddress)
			if err != nil {
				fmt.Println("Error:", err)
				fmt.Println("try 'show ip oui --update' first")
				return
			}
		}
	},
}

func init() {
	ipCmd.AddCommand(ouiCmd)
	ouiCmd.Flags().BoolP("update", "u", false, "update OUI file")
}
