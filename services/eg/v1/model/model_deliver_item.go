package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeliverItem struct {

	// 订阅ID，全局唯一
	SubscriptionId *string `json:"subscriptionId,omitempty"`

	// 订阅名称
	SubscriptionName *string `json:"subscriptionName,omitempty"`

	// 成功目标个数
	SuccessCounts *int32 `json:"successCounts,omitempty"`

	// 失败目标个数
	FailCounts *int32 `json:"failCounts,omitempty"`

	// 共投递次数
	AllDeliverTimes *int32 `json:"allDeliverTimes,omitempty"`

	// 投递详情
	DeliverTargetList *[]DeliverTarget `json:"deliverTargetList,omitempty"`
}

func (o DeliverItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeliverItem struct{}"
	}

	return strings.Join([]string{"DeliverItem", string(data)}, " ")
}
