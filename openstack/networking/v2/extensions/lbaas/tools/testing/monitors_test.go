package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas/monitors"
	th "github.com/gophercloud/gophercloud/testhelper"

	"github.com/jtopjian/gophertools/openstack/networking/v2/extensions/lbaas/tools"
)

func TestGetMonitorType(t *testing.T) {
	monitor, err := tools.GetMonitorType("tcp")
	th.AssertNoErr(t, err)
	th.AssertEquals(t, monitor, monitors.TypeTCP)

	monitor, err = tools.GetMonitorType("nfs")
	if err == nil {
		t.Fatalf("Expected error")
	}
}
