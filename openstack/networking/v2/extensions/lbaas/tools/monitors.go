package tools

import (
	"fmt"
	"strings"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas/monitors"
)

// GetMonitorType takes a string and returns the monitors.MonitorType type.
func GetMonitorType(v string) (monitors.MonitorType, error) {
	var monitor monitors.MonitorType
	var err error

	switch strings.ToUpper(v) {
	case "HTTP":
		monitor = monitors.TypeHTTP
	case "HTTPS":
		monitor = monitors.TypeHTTPS
	case "PING":
		monitor = monitors.TypePING
	case "TCP":
		monitor = monitors.TypeTCP
	default:
		err = fmt.Errorf("Invalid protocol")
	}

	return monitor, err
}
