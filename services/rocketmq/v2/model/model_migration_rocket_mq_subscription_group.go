package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrationRocketMqSubscriptionGroup **参数解释**： RocketMQ元数据迁移，RocketMQ消费组元数据。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type MigrationRocketMqSubscriptionGroup struct {

	// **参数解释**： 消费组名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**： 是否允许以广播模式消费。 **约束限制**： 不涉及。 **取值范围**： - true：允许以广播模式消费。 - false：不允许以广播模式消费。 **默认取值**： 不涉及。
	ConsumeBroadcastEnable *bool `json:"consume_broadcast_enable,omitempty"`

	// **参数解释**： 是否允许消费。 **约束限制**： 不涉及。 **取值范围**： - true：允许消费。 - false：不允许消费。 **默认取值**： true。
	ConsumeEnable *bool `json:"consume_enable,omitempty"`

	// **参数解释**： 是否从最小偏移量开始消费。 **约束限制**： 不涉及。 **取值范围**： - true：是。 - false：不是。 **默认取值**： true。
	ConsumeFromMinEnable *bool `json:"consume_from_min_enable,omitempty"`

	// **参数解释**： 消费者ID变化时是否通知。 **约束限制**： 不涉及。 **取值范围**： - true：是。 - false：不是。 **默认取值**： true。
	NotifyConsumerIdsChangedEnable *bool `json:"notify_consumer_ids_changed_enable,omitempty"`

	// **参数解释**： 消费最大重试次数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 16。
	RetryMaxTimes *int32 `json:"retry_max_times,omitempty"`

	// **参数解释**： 消费最大重试次数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 1。
	RetryQueueNums *int32 `json:"retry_queue_nums,omitempty"`

	// **参数解释**： 慢消费时选择的broker节点ID **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 1。
	WhichBrokerWhenConsumeSlow *int64 `json:"which_broker_when_consume_slow,omitempty"`
}

func (o MigrationRocketMqSubscriptionGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationRocketMqSubscriptionGroup struct{}"
	}

	return strings.Join([]string{"MigrationRocketMqSubscriptionGroup", string(data)}, " ")
}
