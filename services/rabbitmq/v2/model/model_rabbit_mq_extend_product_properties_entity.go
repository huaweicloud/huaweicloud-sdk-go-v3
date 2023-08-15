package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RabbitMqExtendProductPropertiesEntity 功能特性的键值对。
type RabbitMqExtendProductPropertiesEntity struct {

	// Broker的最大个数。
	MaxBroker *string `json:"max_broker,omitempty"`

	// 每个节点的最大存储。单位为GB。
	MaxStoragePerNode *string `json:"max_storage_per_node,omitempty"`

	// Broker的最小个数。
	MinBroker *string `json:"min_broker,omitempty"`

	// 每个节点的最小存储。单位为GB。
	MinStoragePerNode *string `json:"min_storage_per_node,omitempty"`

	// 最大连接数
	MaxConnectionPerBroker *string `json:"max_connection_per_broker,omitempty"`

	// 步长
	StepLength *string `json:"step_length,omitempty"`

	// product_id的别名。
	ProductAlias *string `json:"product_alias,omitempty"`

	// 最大队列
	MaxQueuePerBroker *string `json:"max_queue_per_broker,omitempty"`
}

func (o RabbitMqExtendProductPropertiesEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RabbitMqExtendProductPropertiesEntity struct{}"
	}

	return strings.Join([]string{"RabbitMqExtendProductPropertiesEntity", string(data)}, " ")
}
