package tools

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"

	"github.com/jtopjian/gophertools"
)

// RefreshServerStatus polls a server for its status.
func RefreshServerStatus(client *gophercloud.ServiceClient, serverID string) gophertools.ResourceRefreshFunc {
	return func() (interface{}, string, error) {
		latest, err := servers.Get(client, serverID).Extract()
		if err != nil {
			if _, ok := err.(gophercloud.ErrDefault404); ok {
				return latest, "DELETED", nil
			}
			return nil, "", err
		}

		return latest, latest.Status, nil
	}
}
