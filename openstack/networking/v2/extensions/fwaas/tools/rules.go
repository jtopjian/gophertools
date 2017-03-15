package tools

import (
	"fmt"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas/rules"
)

// GetRuleProtocol takes a string and returns the rules.Protocol type.
func GetRuleProtocol(v string) (rules.Protocol, error) {
	var protocol rules.Protocol
	var err error

	switch v {
	case "any":
		protocol = rules.ProtocolAny
	case "icmp":
		protocol = rules.ProtocolICMP
	case "tcp":
		protocol = rules.ProtocolTCP
	case "udp":
		protocol = rules.ProtocolUDP
	default:
		err = fmt.Errorf("Invalid protocol")
	}

	return protocol, err
}
