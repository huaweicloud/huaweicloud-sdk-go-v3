package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type AclPolicyOption struct {
	// 允许访问的IP地址或网段。

	AllowAddressNetmasks []AllowAddressNetmasksOption `json:"allow_address_netmasks"`
	// 允许访问的IP地址区间。

	AllowIpRanges []AllowIpRangesOption `json:"allow_ip_ranges"`
}

func (o AclPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AclPolicyOption struct{}"
	}

	return strings.Join([]string{"AclPolicyOption", string(data)}, " ")
}
