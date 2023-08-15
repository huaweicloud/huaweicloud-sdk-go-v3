package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApisBindedToRequestThrottlingPolicyV2Request Request Object
type ListApisBindedToRequestThrottlingPolicyV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量，条目数量小于等于0时，自动转换为20，条目数量大于500时，自动转换为500
	Limit *int32 `json:"limit,omitempty"`

	// 流控策略编号
	ThrottleId string `json:"throttle_id"`

	// 环境的ID
	EnvId *string `json:"env_id,omitempty"`

	// API分组编号
	GroupId *string `json:"group_id,omitempty"`

	// API编号
	ApiId *string `json:"api_id,omitempty"`

	// API名称
	ApiName *string `json:"api_name,omitempty"`
}

func (o ListApisBindedToRequestThrottlingPolicyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApisBindedToRequestThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"ListApisBindedToRequestThrottlingPolicyV2Request", string(data)}, " ")
}
