package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas/pools"
	th "github.com/gophercloud/gophercloud/testhelper"

	"github.com/jtopjian/gophertools/openstack/networking/v2/extensions/lbaas/tools"
)

func TestGetPoolProtocol(t *testing.T) {
	proto, err := tools.GetPoolProtocol("tcp")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, proto, pools.ProtocolTCP)

	proto, err = tools.GetPoolProtocol("nfs")
	if err == nil {
		t.Fatalf("Expected error")
	}
}

func TestGetLoadBalancerMethod(t *testing.T) {
	lbm, err := tools.GetLoadBalancerMethod("round_robin")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, lbm, pools.LBMethodRoundRobin)

	lbm, err = tools.GetLoadBalancerMethod("most_connections")
	if err == nil {
		t.Fatalf("Expected error")
	}
}
