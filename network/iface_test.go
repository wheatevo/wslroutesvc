package network_test

import (
	"fmt"
	"testing"

	"github.com/wheatevo/wslroutesvc/network"
)

type missingNewIfaceRunner struct{}

func (e *missingNewIfaceRunner) Run(name string, arg ...string) ([]byte, error) {
	return []byte(""), fmt.Errorf("Could not run command")
}

type foundNewIfaceRunner struct {
	cmdCount int
}

const netshIfaceOutput = `
Interface vEthernet (WSL) Parameters
----------------------------------------------
IfLuid                             : ethernet_32775
IfIndex                            : 39
State                              : connected
Metric                             : 15
Link MTU                           : 1500 bytes
Reachable Time                     : 30500 ms
Base Reachable Time                : 30000 ms
Retransmission Interval            : 1000 ms
DAD Transmits                      : 3
Site Prefix Length                 : 64
Site Id                            : 1
Forwarding                         : disabled
Advertising                        : disabled
Neighbor Discovery                 : enabled
Neighbor Unreachability Detection  : enabled
Router Discovery                   : dhcp
Managed Address Configuration      : enabled
Other Stateful Configuration       : enabled
Weak Host Sends                    : disabled
Weak Host Receives                 : disabled
Use Automatic Metric               : enabled
Ignore Default Routes              : disabled
Advertised Router Lifetime         : 1800 seconds
Advertise Default Route            : disabled
Current Hop Limit                  : 0
Force ARPND Wake up patterns       : disabled
Directed MAC Wake up patterns      : disabled
ECN capability                     : application
RA Based DNS Config (RFC 6106)     : disabled
DHCP/Static IP coexistence         : disabled

`

const netshConfigOutput = `
Configuration for interface "vEthernet (WSL)"
    DHCP enabled:                         No
    IP Address:                           172.29.192.1
    Subnet Prefix:                        172.29.192.0/20 (mask 255.255.240.0)
    InterfaceMetric:                      15
    Statically Configured DNS Servers:    None
    Register with which suffix:           None
    Statically Configured WINS Servers:   None
`

func (e *foundNewIfaceRunner) Run(name string, arg ...string) ([]byte, error) {
	response := ""

	if e.cmdCount == 0 {
		response = netshIfaceOutput
	} else if e.cmdCount == 1 {
		response = netshConfigOutput
	}

	e.cmdCount++
	return []byte(response), nil
}

func TestNewIface(t *testing.T) {
	// when iface cannot be found
	iface := network.NewIface("some missing interface", &missingNewIfaceRunner{})

	if iface.ID != "" || iface.IP.String() != "<nil>" || iface.Network.String() != "<nil>" {
		t.Errorf("Expected interface to have empty values, received %v", iface)
	}

	// when iface is found
	iface = network.NewIface("vEthernet (WSL)", &foundNewIfaceRunner{0})

	if iface.ID != "39" || iface.IP.String() != "172.29.192.1" || iface.Network.String() != "172.29.192.0/20" {
		t.Errorf("Expected interface to have populated values, received %v", iface)
	}
}
