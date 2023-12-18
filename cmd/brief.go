package cmd

import (
	"fmt"
	"log"

	"github.com/bi-zone/wmi"
	"github.com/spf13/cobra"
)

type brief_win32_NetworkAdapterConfiguration struct {
	IPAddress   []string
	IPSubnet    []string
	IPEnabled   bool
	MACAddress  string
	Description string
}

// briefCmd represents the brief command
var briefCmd = &cobra.Command{
	Use:   "brief",
	Short: "Get a brief interface related information",
	Long:  `Get a brief interface related information`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("brief called")
		var dst []brief_win32_NetworkAdapterConfiguration

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
	interfaceCmd.AddCommand(briefCmd)
}
