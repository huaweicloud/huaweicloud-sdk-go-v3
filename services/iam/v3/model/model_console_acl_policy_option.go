package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConsoleAclPolicyOption
type ConsoleAclPolicyOption struct {

	// 允许访问的IP地址或网段。
	AllowAddressNetmasks *[]AllowAddressNetmasksOption `json:"allow_address_netmasks,omitempty"`

	// 允许访问的IP地址区间。
	AllowIpRanges *[]AllowIpRangesOption `json:"allow_ip_ranges,omitempty"`

	// 允许访问的IPv6地址或网段。
	AllowAddressNetmasksIpv6 *[]AllowAddressNetmasksOptionIPv6 `json:"allow_address_netmasks_ipv6,omitempty"`

	// 允许访问的IPv6地址区间。
	AllowIpRangesIpv6 *[]AllowIpRangesOptionIPv6 `json:"allow_ip_ranges_ipv6,omitempty"`
}

func (o ConsoleAclPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsoleAclPolicyOption struct{}"
	}

	return strings.Join([]string{"ConsoleAclPolicyOption", string(data)}, " ")
}
