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
		fmt.Println("IP Address\tSubnet Mask\tMAC Address\t\tIP Enabled\tInterface Description")
		for _, v := range dst {
			if v.IPEnabled {
				fmt.Printf("%v\t%v\t%v\t%v\t\t%v\n", v.IPAddress[0], v.IPSubnet[0], v.MACAddress, v.IPEnabled, v.Description)
			}
		}
	},
}

func init() {
	ipCmd.AddCommand(interfaceCmd)
}
