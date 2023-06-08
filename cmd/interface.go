package cmd

import (
	"fmt"
	"log"

	"github.com/bi-zone/wmi"
	"github.com/spf13/cobra"
)

type interface_win32_NetworkAdapterConfiguration struct {
	IPAddress   []string
	IPSubnet    []string
	IPEnabled   bool
	MACAddress  string
	Description string
}

// interfaceCmd represents the interface command
var interfaceCmd = &cobra.Command{
	Use:   "interface",
	Short: "get interface related information",
	Long:  `get interface related information`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("interface called")
		var dst []interface_win32_NetworkAdapterConfiguration

		q := wmi.CreateQueryFrom(&dst, "Win32_NetworkAdapterConfiguration", "")
		if err := wmi.Query(q, &dst); err != nil {
			log.Fatal(err)
		}
		for _, v := range dst {
			if v.IPEnabled {
				fmt.Print("\n")
				fmt.Printf("Name       : %v\n", v.Description)
				fmt.Printf("IP Address : %v\n", v.IPAddress[0])
				fmt.Printf("Subnet Mask: %v\n", v.IPSubnet[0])
				fmt.Printf("MAC Address: %v\n", v.MACAddress)
				fmt.Printf("Active     : %v\n", v.IPEnabled)
				fmt.Print("\n")
			}
		}
		// }
	},
}

func init() {
	ipCmd.AddCommand(interfaceCmd)
}
