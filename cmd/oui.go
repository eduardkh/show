/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

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
	appDataDir = filepath.Join(appDataDir, "MyOUIApp")
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

// ouiCmd represents the oui command
var ouiCmd = &cobra.Command{
	Use:   "oui",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		update, _ := cmd.Flags().GetBool("update")
		if update {
			err := DownloadOUIFile()
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("oui called")
		}
	},
}

func init() {
	ipCmd.AddCommand(ouiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ouiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	ouiCmd.Flags().BoolP("update", "u", false, "update OUI file")
}
