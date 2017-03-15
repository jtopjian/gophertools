package tools

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/pools"
	"github.com/jtopjian/gophertools"
)

// DeleteMember will attempt to delete a member from a pool.
// It accounts for the various response codes that the Networking API can return.
func DeleteMember(client *gophercloud.ServiceClient, poolID string, memberID string) gophertools.ResourceRefreshFunc {
	return func() (interface{}, string, error) {
		member, err := pools.GetMember(client, poolID, memberID).Extract()
		if err != nil {
			if _, ok := err.(gophercloud.ErrDefault404); ok {
				return member, "DELETED", nil
			}
			return member, "ACTIVE", err
		}

		err = pools.DeleteMember(client, poolID, memberID).ExtractErr()
		if err != nil {
			switch errCode := err.(type) {
			case gophercloud.ErrDefault404:
				return member, "DELETED", nil
			case gophercloud.ErrDefault500:
				return member, "PENDING_DELETE", nil
			case gophercloud.ErrUnexpectedResponseCode:
				if errCode.Actual == 409 {
					return member, "PENDING_DELETE", nil
				}
			default:
				return member, "ACTIVE", err
			}
		}

		return member, "ACTIVE", nil
	}
}
