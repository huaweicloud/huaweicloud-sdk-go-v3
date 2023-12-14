package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrationRabbitExchangeMetadata RocketMQ元数据迁移，RabbitMQ交换机元数据。
type MigrationRabbitExchangeMetadata struct {

	// vhost名称。
	Vhost *string `json:"vhost,omitempty"`

	// 交换机名称。
	Name *string `json:"name,omitempty"`

	// 交换机类型。
	Type *string `json:"type,omitempty"`

	// 是否持久化。
	Durable *bool `json:"durable,omitempty"`
}

func (o MigrationRabbitExchangeMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationRabbitExchangeMetadata struct{}"
	}

	return strings.Join([]string{"MigrationRabbitExchangeMetadata", string(data)}, " ")
}
