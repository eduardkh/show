package cmd

import (
	"fmt"
	"net"
	"strings"

	"github.com/spf13/cobra"
)

// ip calculator function
func calculateIPDetails(input string) {
	var ipnet *net.IPNet
	var ip net.IP

	if strings.Contains(input, "/") {
		// Parse CIDR notation
		var err error
		ip, ipnet, err = net.ParseCIDR(input)
		if err != nil {
			fmt.Println("Invalid input. Please provide a valid CIDR notation.")
			return
		}
	} else {
		// Parse IP and subnet mask
		splitInput := strings.Split(input, " ")
		ip = net.ParseIP(splitInput[0])
		if ip == nil {
			fmt.Println("Invalid IP address. Please provide a valid IP address.")
			return
		}
		mask := net.IPMask(net.ParseIP(splitInput[1]).To4())
		if mask == nil {
			fmt.Println("Invalid subnet mask. Please provide a valid subnet mask.")
			return
		}
		ipnet = &net.IPNet{
			IP:   ip,
			Mask: mask,
		}
	}

	// Calculate other details
	network := ipnet.IP.Mask(ipnet.Mask)
	broadcast := make(net.IP, len(network))
	for i := range network {
		broadcast[i] = network[i] | ^ipnet.Mask[i]
	}

	ones, bits := ipnet.Mask.Size()

	var hostmin net.IP
	var hostmax net.IP
	var hosts int
	var message string = ""

	// Check for /31 and /32 subnets (RFC 3021)
	if ones == 31 {
		hostmin = network
		hostmax = broadcast
		hosts = 2
		message = "Special    : P2P Network RFC 3021"
	} else if ones == 32 {
		hostmin = network
		hostmax = broadcast
		hosts = 1
		message = "Special    : Single Host Address"
	} else {
		hostmin = append(network[:len(network)-1], network[len(network)-1]+1)
		hostmax = append(broadcast[:len(broadcast)-1], broadcast[len(broadcast)-1]-1)
		hosts = (1 << (bits - ones)) - 2
	}

	// Print details
	fmt.Println("Address    :", ip.String()+"/"+fmt.Sprint(ones))
	fmt.Println("SubnetMask :", net.IP(ipnet.Mask).String())
	if ones < 30 {
		fmt.Println("Network    :", network.String())
		fmt.Println("Broadcast  :", broadcast.String())
	}
	if ones == 30 {
		message = "Special    : P2P Network"
	}
	fmt.Println("Host Range :", hostmin.String()+" - "+hostmax.String())
	fmt.Println("Host Number:", hosts)
	fmt.Println(message)
}

// calcCmd represents the calc command
var calcCmd = &cobra.Command{
	Use:   "calc",
	Short: "IP Calculator",
	Long: `IP Calculator
example:
show ip calc 192.168.1.1/25
show ip calc "192.168.1.1 255.255.255.224"`,
	Run: func(cmd *cobra.Command, args []string) {
		ipAddresses := args[0]
		// fmt.Println(ipAddresses)
		//
		// code goes here
		calculateIPDetails(ipAddresses)
		//
	},
}

func init() {
	ipCmd.AddCommand(calcCmd)
}
