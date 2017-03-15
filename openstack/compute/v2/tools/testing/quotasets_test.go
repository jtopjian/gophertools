package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/quotasets"
	th "github.com/gophercloud/gophercloud/testhelper"

	"github.com/jtopjian/gophertools/openstack/compute/v2/tools"
)

func TestCopyQuotaSet(t *testing.T) {
	destQuotaSet := quotasets.UpdateOpts{}
	tools.CopyQuotaSet(srcQuotaSet, &destQuotaSet)

	th.AssertEquals(t, *destQuotaSet.Ram, 1024)
}
