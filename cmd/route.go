/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/bi-zone/wmi"

	"github.com/spf13/cobra"
)

type win32_IP4RouteTable struct {
	Destination string
	Mask        string
	NextHop     string
}

// routeCmd represents the route command
var routeCmd = &cobra.Command{
	Use:   "route",
	Short: "Display IPv4 routing table",
	Long:  `Displays the IPv4 routing table entries, showing the destination, mask, and next hop for the default gateway.`,
	Run: func(cmd *cobra.Command, args []string) {
		var routeTable []win32_IP4RouteTable
		query := wmi.CreateQueryFrom(&routeTable, "Win32_IP4RouteTable", "")
		if err := wmi.Query(query, &routeTable); err != nil {
			log.Fatal(err)
		}

		for _, route := range routeTable {
			if strings.HasPrefix(route.Destination, "224.0.0.") {
				// Skip multicast addresses
				continue
			}
			if route.Destination == "0.0.0.0" && route.Mask == "0.0.0.0" {
				// Displaying only the default gateway
				fmt.Printf("* %s %s via: %s (Default Gateway)\n", route.Destination, route.Mask, route.NextHop)
			} else if route.Mask != "255.255.255.255" { // Filter out specific masks if needed
				fmt.Printf("%s %s via: %s\n", route.Destination, route.Mask, route.NextHop)
			}
		}
	},
}

func init() {
	ipCmd.AddCommand(routeCmd)
}
