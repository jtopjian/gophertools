package testing

import (
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/quotasets"
)

var srcQuotaSet = quotasets.QuotaSet{
	FixedIps:                 4,
	FloatingIps:              1,
	InjectedFileContentBytes: 1024,
	InjectedFilePathBytes:    1024,
	InjectedFiles:            10,
	KeyPairs:                 5,
	Ram:                      1024,
	SecurityGroupRules:       100,
	SecurityGroups:           100,
	Cores:                    4,
	Instances:                4,
	ServerGroups:             2,
	ServerGroupMembers:       4,
	MetadataItems:            100,
}
