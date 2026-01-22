package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrationRabbitExchangeMetadata **参数解释**： RocketMQ元数据迁移，RabbitMQ交换机元数据。 **约束限制**： 不涉及。 **取值范围**： - true：持久化。 - false：不进行持久化。 **默认取值**： 不涉及。
type MigrationRabbitExchangeMetadata struct {

	// **参数解释**： vhost名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Vhost *string `json:"vhost,omitempty"`

	// **参数解释**： 交换机名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 交换机类型。 **约束限制**： 不涉及。 **取值范围**： - topic：支持routing key的模糊匹配。 - direct：按routing key精确匹配进行消息路由。 - fanout：广播模式，消息会发送到所有绑定的队列，忽略routing key。 - headers：根据消息头（headers）中的键值对进行路由，而不是routing key。 **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 是否持久化。 **约束限制**： 不涉及。 **取值范围**： - true：持久化。 - false：不进行持久化。 **默认取值**： 不涉及。
	Durable *bool `json:"durable,omitempty"`
}

func (o MigrationRabbitExchangeMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationRabbitExchangeMetadata struct{}"
	}

	return strings.Join([]string{"MigrationRabbitExchangeMetadata", string(data)}, " ")
}
