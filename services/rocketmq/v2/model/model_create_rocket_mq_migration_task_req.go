package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRocketMqMigrationTaskReq struct {

	// **参数解释**： RocketMQ Topic 元数据，键为Topic名，值为topic配置，迁移任务类型为自建RocketMQ上云(rocketmq)时必填。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TopicConfigTable map[string]MigrationRocketMqTopicConfig `json:"topic_config_table,omitempty"`

	// **参数解释**： RocketMQ消费组元数据，键为消费组名，值为消费组配置，迁移任务类型为自建RocketMQ上云(rocketmq)时必填。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SubscriptionGroupTable map[string]MigrationRocketMqSubscriptionGroup `json:"subscription_group_table,omitempty"`

	// **参数解释**： RabbitMQ vhost元数据列表，迁移任务类型为自建RabbitMQ上云(rabbitToRocket)时必填。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Vhosts *[]MigrationRabbitVhostMetadata `json:"vhosts,omitempty"`

	// **参数解释**： RabbitMQ队列元数据列表，迁移任务类型为自建RabbitMQ上云(rabbitToRocket)时必填。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Queues *[]MigrationRabbitQueueMetadata `json:"queues,omitempty"`

	// **参数解释**： RabbitMQ交换机元数据列表，迁移任务类型为自建RabbitMQ上云(rabbitToRocket)时必填。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Exchanges *[]MigrationRabbitExchangeMetadata `json:"exchanges,omitempty"`

	// **参数解释**： RabbitMQ binding元数据列表，迁移任务类型为自建RabbitMQ上云(rabbitToRocket)时必填。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Bindings *[]MigrationRabbitBindingMetadata `json:"bindings,omitempty"`
}

func (o CreateRocketMqMigrationTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRocketMqMigrationTaskReq struct{}"
	}

	return strings.Join([]string{"CreateRocketMqMigrationTaskReq", string(data)}, " ")
}
