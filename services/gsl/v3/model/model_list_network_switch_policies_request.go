package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNetworkSwitchPoliciesRequest Request Object
type ListNetworkSwitchPoliciesRequest struct {

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 三网卡版本信息
	Version *int32 `json:"version,omitempty"`

	// 分页查询时每页显示的记录数，默认值为10，取值范围为10-500的整数
	Limit *int64 `json:"limit,omitempty"`

	// 分页查询时的页码数，默认值为1，取值范围为1-1000000的整数
	Offset *int64 `json:"offset,omitempty"`
}

func (o ListNetworkSwitchPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNetworkSwitchPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListNetworkSwitchPoliciesRequest", string(data)}, " ")
}
