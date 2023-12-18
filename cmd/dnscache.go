package cmd

import (
	"log"

	"github.com/bi-zone/wmi"
	"github.com/spf13/cobra"
)

// DNSClientCache represents a DNS client cache record
// https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/hh872334(v=vs.85)
type DNSClientCache struct {
	Name       string
	Entry      string
	Data       string
	Type       uint16
	TimeToLive uint32
}

func queryDNSCache() {
	var dnsCacheEntries []DNSClientCache
	query := "SELECT Name, Data, Type, TimeToLive FROM MSFT_DNSClientCache"
	err := wmi.QueryNamespace(query, &dnsCacheEntries, "root/StandardCimv2")
	if err != nil {
		log.Fatalf("Cimv2 query failed: %v", err)
	}

	for _, entry := range dnsCacheEntries {
		if (entry.Type == 1 || entry.Type == 28) && entry.Data != "" {
			log.Printf(`%s
Entry (A / AAAA)       : %s
Data (IP)              : %s
TTL (Seconds)          : %d

`, entry.Name, entry.Entry, entry.Data, entry.TimeToLive)
		}
	}
}

// dnscacheCmd represents the dnscache command
var dnscacheCmd = &cobra.Command{
	Use:   "dnscache",
	Short: "Represents a record in a DNS client cache.",
	Long:  `Displays active DNS cache entries from the Windows DNS client. This command retrieves records with their names, IP addresses, and TTL values, focusing on A (IPv4) and AAAA (IPv6) records. Useful for quick DNS cache inspection and network troubleshooting.`,

	Run: func(cmd *cobra.Command, args []string) {
		queryDNSCache()
	},
}

func init() {
	rootCmd.AddCommand(dnscacheCmd)
}
