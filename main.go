// +build windows

// Windows service for fixing WSL/VPN Routing conflicts automatically
//
// This service searches for routing conflicts caused by a VPN and
// removes any routes that overlap with the WSL network interface
// to avoid WSL networking failure.
//
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/sys/windows/svc"
)

func usage(errmsg string) {
	fmt.Fprintf(os.Stderr,
		"%s\n\n"+
			"Usage: %s <command>\n"+
			"       Where <command> must be\n"+
			"       install, remove, debug, start, stop, pause, or continue.\n",
		errmsg, os.Args[0])
	os.Exit(2)
}

func main() {
	const svcName = "wslroutesvc"

	isService, err := svc.IsWindowsService()
	if err != nil {
		log.Fatalf("Not currently running in Windows service: %v", err)
	}
	if isService {
		runService(svcName, false)
		return
	}

	if len(os.Args) < 2 {
		usage("No command specified")
	}

	cmd := strings.ToLower(os.Args[1])
	switch cmd {
	case "debug":
		runService(svcName, true)
		return
	case "install":
		err = installService(svcName, "WSL VPN Routing Conflict Service")
	case "remove":
		err = removeService(svcName)
	case "start":
		err = startService(svcName)
	case "stop":
		err = controlService(svcName, svc.Stop, svc.Stopped)
	case "pause":
		err = controlService(svcName, svc.Pause, svc.Paused)
	case "continue":
		err = controlService(svcName, svc.Continue, svc.Running)
	default:
		usage(fmt.Sprintf("Invalid command %s", cmd))
	}
	if err != nil {
		log.Fatalf("Failed to %s %s: %v", cmd, svcName, err)
	}
	return
}
