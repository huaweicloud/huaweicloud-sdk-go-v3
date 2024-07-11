package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRocketMqMigrationTaskReq struct {

	// RocketMQ Topic 元数据，键为Topic名，值为topic配置，迁移任务类型为自建RocketMQ上云(rocketmq)时必填。
	TopicConfigTable map[string]MigrationRocketMqTopicConfig `json:"topicConfigTable,omitempty"`

	// RocketMQ消费组元数据，键为消费组名，值为消费组配置，迁移任务类型为自建RocketMQ上云(rocketmq)时必填。
	SubscriptionGroupTable map[string]MigrationRocketMqSubscriptionGroup `json:"subscriptionGroupTable,omitempty"`

	// RabbitMQ vhost元数据列表，迁移任务类型为自建RabbitMQ上云(rabbitToRocket)时必填。
	Vhosts *[]MigrationRabbitVhostMetadata `json:"vhosts,omitempty"`

	// RabbitMQ队列元数据列表，迁移任务类型为自建RabbitMQ上云(rabbitToRocket)时必填。
	Queues *[]MigrationRabbitQueueMetadata `json:"queues,omitempty"`

	// RabbitMQ交换机元数据列表，迁移任务类型为自建RabbitMQ上云(rabbitToRocket)时必填。
	Exchanges *[]MigrationRabbitExchangeMetadata `json:"exchanges,omitempty"`

	// RabbitMQ binding元数据列表，迁移任务类型为自建RabbitMQ上云(rabbitToRocket)时必填。
	Bindings *[]MigrationRabbitBindingMetadata `json:"bindings,omitempty"`
}

func (o CreateRocketMqMigrationTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRocketMqMigrationTaskReq struct{}"
	}

	return strings.Join([]string{"CreateRocketMqMigrationTaskReq", string(data)}, " ")
}
