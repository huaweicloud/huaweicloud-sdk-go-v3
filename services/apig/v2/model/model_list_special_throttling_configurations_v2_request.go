package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListSpecialThrottlingConfigurationsV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 流控策略的ID

	ThrottleId string `json:"throttle_id"`
	// 特殊流控类型：APP, USER

	ObjectType *string `json:"object_type,omitempty"`
	// 筛选的特殊应用名称

	AppName *string `json:"app_name,omitempty"`
	// 筛选的特殊用户名称

	User *string `json:"user,omitempty"`
	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0

	Offset *int64 `json:"offset,omitempty"`
	// 每页显示的条目数量

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSpecialThrottlingConfigurationsV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSpecialThrottlingConfigurationsV2Request struct{}"
	}

	return strings.Join([]string{"ListSpecialThrottlingConfigurationsV2Request", string(data)}, " ")
}
