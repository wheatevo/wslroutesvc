package network

import (
	"fmt"
	"net"
	"regexp"

	"github.com/wheatevo/wslroutesvc/runner"
)

// Iface describes a Windows network interface
type Iface struct {
	Name    string
	ID      string
	IP      net.IP
	Network *net.IPNet
}

// NewIface creates a new network interface object
func NewIface(name string, runner runner.Runner) Iface {
	n := Iface{name, "", net.IP{}, &net.IPNet{}}

	n.retrieveID(runner)
	n.retrieveIP(runner)

	return n
}

// RetrieveID gets the network interface's ID from netsh
func (n *Iface) retrieveID(runner runner.Runner) error {
	out, err := runner.Run("netsh", "interface", "ipv4", "show", "interfaces", n.Name)

	if err != nil {
		return err
	}

	r := regexp.MustCompile(`IfIndex.*:\s+(\w+)`)

	matches := r.FindStringSubmatch(string(out))

	if len(matches) > 0 {
		n.ID = matches[1]
		return nil
	}

	return fmt.Errorf("could not find interface ID for interface %s", n.Name)
}

// RetrieveIP gets the network interface's IP from netsh
func (n *Iface) retrieveIP(runner runner.Runner) error {
	out, err := runner.Run("netsh", "interface", "ipv4", "show", "config", n.Name)

	if err != nil {
		return err
	}

	r := regexp.MustCompile(`IP Address:\s+([\w\.]+)`)
	matches := r.FindStringSubmatch(string(out))

	if len(matches) > 0 {
		n.IP = net.ParseIP(matches[1])
	}

	r = regexp.MustCompile(`Subnet Prefix:\s+([\w\.\/]+)`)
	matches = r.FindStringSubmatch(string(out))

	if len(matches) > 0 {
		_, network, _ := net.ParseCIDR(matches[1])

		n.Network = network

		return nil
	}

	return fmt.Errorf("could not find interface IP for interface %s", n.Name)
}

func (n Iface) String() string {
	return fmt.Sprintf("%s (ID: %s, IP: %s, Network: %s)", n.Name, n.ID, n.IP, n.Network)
}
