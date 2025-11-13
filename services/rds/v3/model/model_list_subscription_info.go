package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubscriptionInfo 订阅详情。
type ListSubscriptionInfo struct {

	// 订阅id。
	Id *string `json:"id,omitempty"`

	// 订阅状态。normal，abnormal，creating，createfail
	Status *string `json:"status,omitempty"`

	// 发布id。
	PublicationId *string `json:"publication_id,omitempty"`

	// 发布名称。
	PublicationName *string `json:"publication_name,omitempty"`

	// 订阅服务器来源。true：订阅服务器为云上实例, false：订阅服务器非云上实例。
	IsCloud *bool `json:"is_cloud,omitempty"`

	// 目标数据库名。
	SubscriptionDatabase *string `json:"subscription_database,omitempty"`

	// 订阅方式，push:推送。
	SubscriptionType *string `json:"subscription_type,omitempty"`

	PublicationSubscription *Subscription4PublicationInfo `json:"publication_subscription,omitempty"`

	LocalSubscription *Subscription4InstanceInfo `json:"local_subscription,omitempty"`

	JobSchedule *UsedJobSchedule `json:"job_schedule,omitempty"`
}

func (o ListSubscriptionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionInfo struct{}"
	}

	return strings.Join([]string{"ListSubscriptionInfo", string(data)}, " ")
}
