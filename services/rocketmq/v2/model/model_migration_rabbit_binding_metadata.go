package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrationRabbitBindingMetadata **参数解释**： RocketMQ元数据迁移，RabbitMQ binding元数据。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type MigrationRabbitBindingMetadata struct {

	// **参数解释**： vhost名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Vhost *string `json:"vhost,omitempty"`

	// **参数解释**： 消息的来源。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Source *string `json:"source,omitempty"`

	// **参数解释**： 消息的目标。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Destination *string `json:"destination,omitempty"`

	// **参数解释**： 目标的类型。 **约束限制**： 不涉及。 **取值范围**： - exchange：交换机。 - queue：队列。[RabbitMQ AMQP版本只支持绑定queue。](tag:hws,hws_hk) **默认取值**： 不涉及。
	DestinationType *string `json:"destination_type,omitempty"`

	// **参数解释**： 路由键。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	RoutingKey *string `json:"routing_key,omitempty"`
}

func (o MigrationRabbitBindingMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationRabbitBindingMetadata struct{}"
	}

	return strings.Join([]string{"MigrationRabbitBindingMetadata", string(data)}, " ")
}
