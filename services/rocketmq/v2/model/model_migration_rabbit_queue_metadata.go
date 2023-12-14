package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrationRabbitQueueMetadata RocketMQ元数据迁移，RabbitMQ队列元数据。
type MigrationRabbitQueueMetadata struct {

	// vhost名称。
	Vhost *string `json:"vhost,omitempty"`

	// 队列名称。
	Name *string `json:"name,omitempty"`

	// 是否持久化。
	Durable *bool `json:"durable,omitempty"`
}

func (o MigrationRabbitQueueMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationRabbitQueueMetadata struct{}"
	}

	return strings.Join([]string{"MigrationRabbitQueueMetadata", string(data)}, " ")
}
