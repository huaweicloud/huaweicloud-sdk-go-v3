package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AclPolicyResult struct {

	// 允许访问的IP地址或网段。
	AllowAddressNetmasks *[]AllowAddressNetmasksResult `json:"allow_address_netmasks,omitempty"`

	// 允许访问的IP地址区间。
	AllowIpRanges *[]AllowIpRangesResult `json:"allow_ip_ranges,omitempty"`

	// 允许访问的VPC端点。
	AllowVpcEndpoints *[]AllowVpcEndpointsResult `json:"allow_vpc_endpoints,omitempty"`
}

func (o AclPolicyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AclPolicyResult struct{}"
	}

	return strings.Join([]string{"AclPolicyResult", string(data)}, " ")
}
