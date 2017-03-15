package tools

import (
	"fmt"
	"strings"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules"
)

// GetRuleEtherType takes a string and returns the rules.EtherType type.
func GetRuleEtherType(v string) (rules.RuleEtherType, error) {
	var et rules.RuleEtherType
	var err error

	switch strings.ToLower(v) {
	case "ipv4":
		et = rules.EtherType4
	case "ipv6":
		et = rules.EtherType6
	default:
		err = fmt.Errorf("Invalid EtherType")
	}

	return et, err
}

// GetRuleDirection takes a string and returns the rules.RuleDirection type.
func GetRuleDirection(v string) (rules.RuleDirection, error) {
	var direction rules.RuleDirection
	var err error

	switch v {
	case "ingress":
		direction = rules.DirIngress
	case "egress":
		direction = rules.DirEgress
	default:
		err = fmt.Errorf("Invalid direction")
	}

	return direction, err
}

// GetRuleProtocol takes a string and returns the rules.RuleProtocol type.
func GetRuleProtocol(v string) (rules.RuleProtocol, error) {
	var protocol rules.RuleProtocol
	var err error

	switch v {
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
