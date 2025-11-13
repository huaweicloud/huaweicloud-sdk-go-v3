package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifySubscriptionsRequestBody 修改订阅详情。
type ModifySubscriptionsRequestBody struct {

	// 修改的订阅id列表。每次修改的订阅必须属于同一实例。
	SubscriptionIds []string `json:"subscription_ids"`

	JobSchedule *OperateUsedJobSchedule `json:"job_schedule"`
}

func (o ModifySubscriptionsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifySubscriptionsRequestBody struct{}"
	}

	return strings.Join([]string{"ModifySubscriptionsRequestBody", string(data)}, " ")
}
