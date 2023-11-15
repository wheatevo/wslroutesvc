//go:build windows

package main

import (
	"fmt"

	"github.com/wheatevo/wslroutesvc/network"
	"github.com/wheatevo/wslroutesvc/runner"
)

func fixRoutes(wslIfaceNames []string, runner runner.Runner) {
	var wslIface network.Iface
	var wslIfaceName string
	foundIface := false

	// Iterate through possible WSL interface names to find a valid interface
	for _, ifaceName := range wslIfaceNames {
		wslIface = network.NewIface(ifaceName, runner)

		if wslIface.ID == "" || wslIface.IP.String() == "<nil>" {
			continue
		}

		wslIfaceName = ifaceName
		foundIface = true
		break
	}

	if !foundIface {
		elog.Error(1, fmt.Sprintf("Could not find interface ID or IP for WSL interfaces ", wslIfaceNames))
		return
	}

	elog.Info(1, fmt.Sprintf("%s interface ID: %s, IP: %s", wslIfaceName, wslIface.ID, wslIface.IP))

	routeList := network.NewRouteList(runner)

	for _, r := range routeList.Routes {
		if r.Network.Contains(wslIface.IP) && r.InterfaceID != wslIface.ID {
			maskSize, _ := r.Network.Mask.Size()

			// Prevent broad routes from qualifying
			if maskSize < 16 {
				continue
			}

			// Remove the route
			out, err := r.Remove(runner)

			if err != nil {
				elog.Error(1, fmt.Sprintf("Failed to remove route %s with interface ID %s!\n%s\n%v", r.Network, r.InterfaceID, out, err))
				continue
			}

			elog.Info(1, fmt.Sprintf("Route %s with interface ID %s removed!", r.Network, r.InterfaceID))
		}
	}
}
