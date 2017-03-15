package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas/rules"
	th "github.com/gophercloud/gophercloud/testhelper"

	"github.com/jtopjian/gophertools/openstack/networking/v2/extensions/fwaas/tools"
)

func TestGetRuleProtocol(t *testing.T) {
	proto, err := tools.GetRuleProtocol("tcp")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, proto, rules.ProtocolTCP)

	proto, err = tools.GetRuleProtocol("nfs")
	if err == nil {
		t.Fatalf("Expected error")
	}
}
