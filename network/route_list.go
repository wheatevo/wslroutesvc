package network

import (
	"net"
	"os/exec"
	"strings"
)

// RouteList describes one or many Windows network route
type RouteList struct {
	Routes []Route
}

// NewRouteList creates a new network route list from current network routes
func NewRouteList() RouteList {
	r := RouteList{[]Route{}}

	// Gather current route output
	out, err := exec.Command("netsh", "interface", "ipv4", "show", "route").Output()

	if err != nil {
		return r
	}

	routeLines := strings.Split(string(out), "\n")

	for _, l := range routeLines {
		if strings.HasPrefix(l, "Publish") || strings.HasPrefix(l, "------") || strings.TrimSpace(l) == "" {
			continue
		}

		routeFields := strings.Fields(l)
		ifaceID := routeFields[4]

		_, prefix, _ := net.ParseCIDR(routeFields[3])

		lineRoute := NewRoute(*prefix, ifaceID)

		r.Routes = append(r.Routes, lineRoute)
	}

	return r
}
