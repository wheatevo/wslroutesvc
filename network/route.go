package network

import (
	"fmt"
	"net"
	"time"

	"github.com/wheatevo/wslroutesvc/runner"
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
func (r *Route) Remove(runner runner.Runner) ([]byte, error) {
	out, err := runner.Run("netsh", "interface", "ipv4", "delete", "route", r.Network.String(), r.InterfaceID)

	// For some reason this requires multiple removals to work with the VPN, attempt a second removal if the first succeeds
	if err == nil {
		time.Sleep(500 * time.Millisecond)
		runner.Run("netsh", "interface", "ipv4", "delete", "route", r.Network.String(), r.InterfaceID)
	}

	return out, err
}

func (r Route) String() string {
	return fmt.Sprintf("Network: %s, InterfaceID: %s", r.Network, r.InterfaceID)
}
