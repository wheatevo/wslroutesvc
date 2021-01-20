package network_test

import (
	"fmt"
	"net"
	"testing"

	"github.com/wheatevo/wslroutesvc/network"
)

func TestNewRoute(t *testing.T) {
	// when arguments are empty
	emptyNet := net.IPNet{}

	r := network.NewRoute(emptyNet, "")
	if r.InterfaceID != "" || r.Network.String() != emptyNet.String() {
		t.Errorf("Expected route to have empty parameters, received %v", r)
	}

	// when arguments are provided
	_, nw, _ := net.ParseCIDR("192.168.0.5/24")
	r = network.NewRoute(*nw, "123")
	if r.InterfaceID != "123" || r.Network.String() != nw.String() {
		t.Errorf("Expected route to have interface ID of 123 and network of 192.168.0.5/24, received %v", r)
	}
}

type failRemoveRunner struct{}

func (e *failRemoveRunner) Run(name string, arg ...string) ([]byte, error) {
	return []byte("Could not remove the route!"), fmt.Errorf("1")
}

type successRemoveRunner struct{}

func (e *successRemoveRunner) Run(name string, arg ...string) ([]byte, error) {
	return []byte("Route removed"), nil
}

func TestRemove(t *testing.T) {
	// when netsh command fails
	_, nw, _ := net.ParseCIDR("192.168.0.5/24")
	r := network.NewRoute(*nw, "123")
	_, err := r.Remove(&failRemoveRunner{})
	if err == nil {
		t.Errorf("Expected failed route.Remove() to return error")
	}

	// when netsh command succeeds
	out, err := r.Remove(&successRemoveRunner{})
	if string(out) != "Route removed" {
		t.Errorf("Expected route.Remove() output to equal Route removed")
	}

	if err != nil {
		t.Errorf("Expected successful route.Remove() to return no error")
	}
}
