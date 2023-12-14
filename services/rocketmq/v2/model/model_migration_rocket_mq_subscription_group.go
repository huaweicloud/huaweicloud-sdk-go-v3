package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrationRocketMqSubscriptionGroup RocketMQ元数据迁移，RocketMQ消费组元数据。
type MigrationRocketMqSubscriptionGroup struct {

	// 消费组名。
	GroupName *string `json:"groupName,omitempty"`

	// 是否允许以广播模式消费。
	ConsumeBroadcastEnable *bool `json:"consumeBroadcastEnable,omitempty"`

	// 是否允许消费。
	ConsumeEnable *bool `json:"consumeEnable,omitempty"`

	// 是否从最小偏移量开始消费。
	ConsumeFromMinEnable *bool `json:"consumeFromMinEnable,omitempty"`

	// 消费者ID变化时是否通知。
	NotifyConsumerIdsChangedEnable *bool `json:"notifyConsumerIdsChangedEnable,omitempty"`

	// 消费最大重试次数。
	RetryMaxTimes *int32 `json:"retryMaxTimes,omitempty"`

	// 重试队列个数。
	RetryQueueNums *int32 `json:"retryQueueNums,omitempty"`

	// 慢消费时选择的broker节点ID。
	WhichBrokerWhenConsumeSlow *int64 `json:"whichBrokerWhenConsumeSlow,omitempty"`
}

func (o MigrationRocketMqSubscriptionGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationRocketMqSubscriptionGroup struct{}"
	}

	return strings.Join([]string{"MigrationRocketMqSubscriptionGroup", string(data)}, " ")
}
