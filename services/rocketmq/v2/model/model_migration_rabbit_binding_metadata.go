package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrationRabbitBindingMetadata RocketMQ元数据迁移，RabbitMQ binding元数据。
type MigrationRabbitBindingMetadata struct {

	// vhost名称。
	Vhost *string `json:"vhost,omitempty"`

	// 消息的来源。
	Source *string `json:"source,omitempty"`

	// 消息的目标。
	Destination *string `json:"destination,omitempty"`

	// 目标的类型。
	DestinationType *string `json:"destination_type,omitempty"`

	// 路由键。
	RoutingKey *string `json:"routing_key,omitempty"`
}

func (o MigrationRabbitBindingMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationRabbitBindingMetadata struct{}"
	}

	return strings.Join([]string{"MigrationRabbitBindingMetadata", string(data)}, " ")
}
