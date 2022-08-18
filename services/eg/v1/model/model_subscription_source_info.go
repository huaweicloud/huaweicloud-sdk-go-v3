package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubscriptionSourceInfo struct {

	// 订阅源ID
	Id *string `json:"id,omitempty"`

	// 订阅的事件源名称
	Name *string `json:"name,omitempty"`

	// 订阅的事件源的提供方类型
	ProviderType *string `json:"provider_type,omitempty"`

	// 订阅的事件源参数列表
	Detail *interface{} `json:"detail,omitempty"`

	// 订阅事件源的匹配过滤规则
	Filter *interface{} `json:"filter,omitempty"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新时间
	UpdatedTime *string `json:"updated_time,omitempty"`
}

func (o SubscriptionSourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionSourceInfo struct{}"
	}

	return strings.Join([]string{"SubscriptionSourceInfo", string(data)}, " ")
}
