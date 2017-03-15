package tools

import (
	"fmt"
	"strings"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas/pools"
)

// GetPoolProtocol takes a string and returns the pools.Protocol type.
func GetPoolProtocol(v string) (pools.LBProtocol, error) {
	var protocol pools.LBProtocol
	var err error

	switch strings.ToUpper(v) {
	case "HTTP":
		protocol = pools.ProtocolHTTP
	case "HTTPS":
		protocol = pools.ProtocolHTTPS
	case "TCP":
		protocol = pools.ProtocolTCP
	default:
		err = fmt.Errorf("Invalid protocol")
	}

	return protocol, err
}

// GetLoadBalancerMethod takes a string and returns the pools.LBMethod type.
func GetLoadBalancerMethod(v string) (pools.LBMethod, error) {
	var lbmethod pools.LBMethod
	var err error

	switch strings.ToUpper(v) {
	case "ROUND_ROBIN":
		lbmethod = pools.LBMethodRoundRobin
	case "LEAST_CONNECTIONS":
		lbmethod = pools.LBMethodLeastConnections
	default:
		err = fmt.Errorf("Invalid protocol")
	}

	return lbmethod, err
}
