package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConsoleAclPolicyResult
type ConsoleAclPolicyResult struct {

	// 允许访问的IP地址或网段。
	AllowAddressNetmasks *[]AllowAddressNetmasksResult `json:"allow_address_netmasks,omitempty"`

	// 允许访问的IP地址区间。
	AllowIpRanges *[]AllowIpRangesResult `json:"allow_ip_ranges,omitempty"`

	// 允许访问的IPv6地址或网段。
	AllowAddressNetmasksIpv6 *[]AllowAddressNetmasksResultIPv6 `json:"allow_address_netmasks_ipv6,omitempty"`

	// 允许访问的IPv6地址区间。
	AllowIpRangesIpv6 *[]AllowIpRangesResultIPv6 `json:"allow_ip_ranges_ipv6,omitempty"`
}

func (o ConsoleAclPolicyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsoleAclPolicyResult struct{}"
	}

	return strings.Join([]string{"ConsoleAclPolicyResult", string(data)}, " ")
}
