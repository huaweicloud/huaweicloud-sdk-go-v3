package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSpecialThrottlingConfigurationsV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 流控策略的编号
	ThrottleId string `json:"throttle_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量，条目数量小于等于0时，自动转换为20，条目数量大于500时，自动转换为500
	Limit *int32 `json:"limit,omitempty"`

	// 特殊流控类型：APP，USER
	ObjectType *string `json:"object_type,omitempty"`

	// 筛选的特殊应用名称
	AppName *string `json:"app_name,omitempty"`

	// 筛选的特殊用户名称
	User *string `json:"user,omitempty"`
}

func (o ListSpecialThrottlingConfigurationsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpecialThrottlingConfigurationsV2Request struct{}"
	}

	return strings.Join([]string{"ListSpecialThrottlingConfigurationsV2Request", string(data)}, " ")
}
