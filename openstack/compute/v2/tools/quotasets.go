package tools

import (
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/quotasets"
)

// CopyQuotaSet will copy the values from a source quotaset to a
// destination quotaset UpdateOpts. This is useful for using the
// source quotaset as a template for a new set.
func CopyQuotaSet(src quotasets.QuotaSet, dest *quotasets.UpdateOpts) {
	dest.FixedIps = &src.FixedIps
	dest.FloatingIps = &src.FloatingIps
	dest.InjectedFileContentBytes = &src.InjectedFileContentBytes
	dest.InjectedFilePathBytes = &src.InjectedFilePathBytes
	dest.InjectedFiles = &src.InjectedFiles
	dest.KeyPairs = &src.KeyPairs
	dest.Ram = &src.Ram
	dest.SecurityGroupRules = &src.SecurityGroupRules
	dest.SecurityGroups = &src.SecurityGroups
	dest.Cores = &src.Cores
	dest.Instances = &src.Instances
	dest.ServerGroups = &src.ServerGroups
	dest.ServerGroupMembers = &src.ServerGroupMembers
	dest.MetadataItems = &src.MetadataItems
}
