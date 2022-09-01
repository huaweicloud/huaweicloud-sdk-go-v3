package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 订阅的事件目标转换规则
type SubscriptionTargetInfoTransform struct {

	// 转换规则类型
	Type *string `json:"type,omitempty" xml:"type"`

	// 转换规则内容
	Value *string `json:"value,omitempty" xml:"value"`

	// 转换规则模板
	Template *string `json:"template,omitempty" xml:"template"`
}

func (o SubscriptionTargetInfoTransform) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionTargetInfoTransform struct{}"
	}

	return strings.Join([]string{"SubscriptionTargetInfoTransform", string(data)}, " ")
}
