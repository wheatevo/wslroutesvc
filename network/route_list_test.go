package network_test

import (
	"testing"

	"github.com/wheatevo/wslroutesvc/network"
)

const netshShowRouteOutput = `
Publish  Type      Met  Prefix                    Idx  Gateway/Interface Name
-------  --------  ---  ------------------------  ---  ------------------------
No       Manual    0    0.0.0.0/0                  16  192.168.1.1
No       System    256  127.0.0.0/8                 1  Loopback Pseudo-Interface 1
No       System    256  127.0.0.1/32                1  Loopback Pseudo-Interface 1
No       System    256  127.255.255.255/32          1  Loopback Pseudo-Interface 1
No       System    256  172.29.192.0/20            39  vEthernet (WSL)
No       System    256  172.29.192.1/32            39  vEthernet (WSL)
No       System    256  172.29.207.255/32          39  vEthernet (WSL)
No       System    256  192.168.1.0/24             16  Wi-Fi
No       System    256  192.168.1.225/32           16  Wi-Fi
No       System    256  192.168.1.255/32           16  Wi-Fi
No       System    256  224.0.0.0/4                 1  Loopback Pseudo-Interface 1
No       System    256  224.0.0.0/4                 8  Ethernet
No       System    256  224.0.0.0/4                16  Wi-Fi
No       System    256  224.0.0.0/4                18  Local Area Connection* 1
No       System    256  224.0.0.0/4                17  Local Area Connection* 10
No       System    256  224.0.0.0/4                39  vEthernet (WSL)
No       System    256  255.255.255.255/32          1  Loopback Pseudo-Interface 1
No       System    256  255.255.255.255/32          8  Ethernet
No       System    256  255.255.255.255/32         16  Wi-Fi
No       System    256  255.255.255.255/32         18  Local Area Connection* 1
No       System    256  255.255.255.255/32         17  Local Area Connection* 10
No       System    256  255.255.255.255/32         39  vEthernet (WSL)
`

type newRouteListRunner struct{}

func (e *newRouteListRunner) Run(name string, arg ...string) ([]byte, error) {
	return []byte(netshShowRouteOutput), nil
}

func TestNewRouteList(t *testing.T) {
	// when routes are found
	routeList := network.NewRouteList(&newRouteListRunner{})
	numRoutes := 22

	if len(routeList.Routes) != numRoutes {
		t.Errorf("Expected %d routes found %d routes.", numRoutes, len(routeList.Routes))
	}

	// spot test first route
	routeOne := routeList.Routes[0]

	if routeOne.InterfaceID != "16" || routeOne.Network.String() != "0.0.0.0/0" {
		t.Errorf("Expected %v to have an interface ID of 16 and a network of 0.0.0.0/0", routeOne)
	}
}
