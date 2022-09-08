package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubscriptionTargetInfo struct {

	// 订阅目标ID
	Id *string `json:"id,omitempty"`

	// 订阅的事件目标名称
	Name *string `json:"name,omitempty"`

	// 订阅的事件目标的提供方类型
	ProviderType *string `json:"provider_type,omitempty"`

	// 订阅的事件目标参数列表
	Detail *interface{} `json:"detail,omitempty"`

	Transform *SubscriptionTargetInfoTransform `json:"transform,omitempty"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新时间
	UpdatedTime *string `json:"updated_time,omitempty"`
}

func (o SubscriptionTargetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionTargetInfo struct{}"
	}

	return strings.Join([]string{"SubscriptionTargetInfo", string(data)}, " ")
}
