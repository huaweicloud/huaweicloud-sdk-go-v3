package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNetworkSwitchPoliciesResponse Response Object
type ListNetworkSwitchPoliciesResponse struct {

	// 每页的记录数
	Limit *int64 `json:"limit,omitempty"`

	// 页码，最小值是1，最大值为1000000。默认值是1.
	Offset *int64 `json:"offset,omitempty"`

	// 记录总数
	Count *int64 `json:"count,omitempty"`

	// 网络切换策略实例列表
	NetworkSwitchPolicyList *[]NetworkSwitchPolicyVo `json:"network_switch_policy_list,omitempty"`
	HttpStatusCode          int                      `json:"-"`
}

func (o ListNetworkSwitchPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNetworkSwitchPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListNetworkSwitchPoliciesResponse", string(data)}, " ")
}
