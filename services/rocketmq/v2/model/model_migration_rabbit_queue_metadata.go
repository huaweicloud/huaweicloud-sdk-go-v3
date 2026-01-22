package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrationRabbitQueueMetadata **参数解释**： RocketMQ元数据迁移，RabbitMQ队列元数据。 **约束限制**： 不涉及。 **取值范围**： - true：持久化。 - false：不进行持久化。 **默认取值**： 不涉及。
type MigrationRabbitQueueMetadata struct {

	// **参数解释**： vhost名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Vhost *string `json:"vhost,omitempty"`

	// **参数解释**： 队列名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 是否持久化。 **约束限制**： 不涉及。 **取值范围**： - true：持久化。 - false：不进行持久化。 **默认取值**： 不涉及。
	Durable *bool `json:"durable,omitempty"`
}

func (o MigrationRabbitQueueMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationRabbitQueueMetadata struct{}"
	}

	return strings.Join([]string{"MigrationRabbitQueueMetadata", string(data)}, " ")
}
