package findnetworkpkg

import (
	"fmt"
	"log"
	"net"
)

func Get() {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatalf("failed to get network interfaces: %v", err)
	}

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			log.Fatalf("failed to get addresses for interface %v: %v", iface.Name, err)
		}

		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			if ipnet.IP.IsLoopback() || ipnet.IP.IsLinkLocalUnicast() {
				continue
			}

			fmt.Printf("Interface: %v, Address: %v\n", iface.Name, addr.String())

			if iface.HardwareAddr.String() != "" {
				fmt.Printf("  Hardware Address (MAC): %v\n", iface.HardwareAddr)
			}

			// Check if the interface is wireless (WiFi)
			if isWirelessInterface(iface.Flags) {
				fmt.Println("  Interface is a WiFi interface")
				// You can add more specific checks or actions here for WiFi interfaces
			}
		}
	}
}

// Check if the interface is a wireless (WiFi) interface based on flags
func isWirelessInterface(flags net.Flags) bool {
	return flags&net.FlagBroadcast != 0 && flags&net.FlagUp != 0
}
