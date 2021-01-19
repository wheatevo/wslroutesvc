package network

import (
	"fmt"
	"net"
	"os/exec"
)

// Route describes a Windows network route
type Route struct {
	Network     net.IPNet
	InterfaceID string
}

// NewRoute creates a new network route object
func NewRoute(network net.IPNet, ifaceID string) Route {
	r := Route{network, ifaceID}
	return r
}

// Remove removes an existing route from the routing table
func (r *Route) Remove() ([]byte, error) {
	out, err := exec.Command("netsh", "interface", "ipv4", "delete", "route", r.Network.String(), r.InterfaceID).Output()

	return out, err
}

func (r Route) String() string {
	return fmt.Sprintf("Network: %s, InterfaceID: %s", r.Network, r.InterfaceID)
}
