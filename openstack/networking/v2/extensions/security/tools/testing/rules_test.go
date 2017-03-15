package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules"
	th "github.com/gophercloud/gophercloud/testhelper"

	"github.com/jtopjian/gophertools/openstack/networking/v2/extensions/security/tools"
)

func TestGetRuleEtherType(t *testing.T) {
	et, err := tools.GetRuleEtherType("ipv4")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, et, rules.EtherType4)

	et, err = tools.GetRuleEtherType("ipv464")
	if err == nil {
		t.Fatalf("Expected error")
	}
}

func TestGetRuleDirection(t *testing.T) {
	d, err := tools.GetRuleDirection("ingress")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, d, rules.DirIngress)

	d, err = tools.GetRuleDirection("upgress")
	if err == nil {
		t.Fatalf("Expected error")
	}
}

func TestGetRuleProtocol(t *testing.T) {
	proto, err := tools.GetRuleProtocol("tcp")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, proto, rules.ProtocolTCP)

	proto, err = tools.GetRuleProtocol("nfs")
	if err == nil {
		t.Fatalf("Expected error")
	}
}
