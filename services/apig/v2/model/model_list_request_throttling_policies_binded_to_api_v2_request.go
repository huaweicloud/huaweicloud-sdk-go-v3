package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListRequestThrottlingPoliciesBindedToApiV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// API编号

	ApiId string `json:"api_id"`
	// 流控策略的编号

	ThrottleId *string `json:"throttle_id,omitempty"`
	// 流控策略的名称

	ThrottleName *string `json:"throttle_name,omitempty"`
	// 绑定的环境编号

	EnvId *string `json:"env_id,omitempty"`
	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0

	Offset *int64 `json:"offset,omitempty"`
	// 每页显示的条目数量

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRequestThrottlingPoliciesBindedToApiV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListRequestThrottlingPoliciesBindedToApiV2Request struct{}"
	}

	return strings.Join([]string{"ListRequestThrottlingPoliciesBindedToApiV2Request", string(data)}, " ")
}
