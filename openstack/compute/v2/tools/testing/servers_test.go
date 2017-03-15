package testing

import (
	"testing"
	"time"

	gt "github.com/gophercloud/gophercloud/openstack/compute/v2/servers/testing"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"

	"github.com/jtopjian/gophertools"
	"github.com/jtopjian/gophertools/openstack/compute/v2/tools"
)

func TestServerStatus(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	gt.HandleServerGetSuccessfully(t)

	rrc := gophertools.ResourceRefreshConf{
		Timeout:      60 * time.Second,
		PollInterval: 1 * time.Second,
		Target:       []string{"ACTIVE"},
		Refresh:      tools.RefreshServerStatus(client.ServiceClient(), "1234asdf"),
	}

	_, err := rrc.ResourceRefresh()
	th.AssertNoErr(t, err)
}
