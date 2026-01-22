package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrationRabbitVhostMetadata **参数解释**： RocketMQ元数据迁移，RabbitMQ vhost元数据。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type MigrationRabbitVhostMetadata struct {

	// **参数解释**： vhost名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o MigrationRabbitVhostMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationRabbitVhostMetadata struct{}"
	}

	return strings.Join([]string{"MigrationRabbitVhostMetadata", string(data)}, " ")
}
