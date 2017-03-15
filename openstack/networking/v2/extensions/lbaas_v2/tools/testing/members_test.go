package testing

import (
	"testing"
	"time"

	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"

	"github.com/jtopjian/gophertools"
	"github.com/jtopjian/gophertools/openstack/networking/v2/extensions/lbaas_v2/tools"
)

func TestMemberDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleMemberDeleteCustomCode(t, 404)

	rrc := gophertools.ResourceRefreshConf{
		Timeout:      60 * time.Second,
		PollInterval: 1 * time.Second,
		Target:       []string{"DELETED"},
		Refresh:      tools.DeleteMember(client.ServiceClient(), "1234asdf", "1234asdf"),
	}

	_, err := rrc.ResourceRefresh()
	th.AssertNoErr(t, err)
}

func TestMemberDelete409(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleMemberDeleteCustomCode(t, 409)

	rrc := gophertools.ResourceRefreshConf{
		Timeout:      60 * time.Second,
		PollInterval: 1 * time.Second,
		Target:       []string{"PENDING_DELETE"},
		Refresh:      tools.DeleteMember(client.ServiceClient(), "1234asdf", "1234asdf"),
	}

	_, err := rrc.ResourceRefresh()
	th.AssertNoErr(t, err)
}

func TestMemberDelete500(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleMemberDeleteCustomCode(t, 500)

	rrc := gophertools.ResourceRefreshConf{
		Timeout:      60 * time.Second,
		PollInterval: 1 * time.Second,
		Target:       []string{"PENDING_DELETE"},
		Refresh:      tools.DeleteMember(client.ServiceClient(), "1234asdf", "1234asdf"),
	}

	_, err := rrc.ResourceRefresh()
	th.AssertNoErr(t, err)
}

func TestMemberDeleteStillActive(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleMemberDeleteCustomCode(t, 200)

	rrc := gophertools.ResourceRefreshConf{
		Timeout:      60 * time.Second,
		PollInterval: 1 * time.Second,
		Target:       []string{"ACTIVE"},
		Refresh:      tools.DeleteMember(client.ServiceClient(), "1234asdf", "1234asdf"),
	}

	_, err := rrc.ResourceRefresh()
	th.AssertNoErr(t, err)
}
