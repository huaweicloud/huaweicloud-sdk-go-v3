package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RocketMqExtendProductPropertiesEntity 功能特性的键值对。
type RocketMqExtendProductPropertiesEntity struct {

	// Broker的最大个数。
	MaxBroker *string `json:"max_broker,omitempty"`

	// 每个节点最多能创建的Topic个数。
	MaxTopicPerBroker *string `json:"max_topic_per_broker,omitempty"`

	// 每个节点的最大消费者数。
	MaxConsumerPerBroker *string `json:"max_consumer_per_broker,omitempty"`

	// 每个节点的最大存储。单位为GB
	MaxStoragePerNode *string `json:"max_storage_per_node,omitempty"`

	// Broker的最小个数。
	MinBroker *string `json:"min_broker,omitempty"`

	// 消息引擎版本。
	EngineVersions *string `json:"engine_versions,omitempty"`

	// 每个节点的最小存储。单位为GB。
	MinStoragePerNode *string `json:"min_storage_per_node,omitempty"`

	// product_id的别名。
	ProductAlias *string `json:"product_alias,omitempty"`
}

func (o RocketMqExtendProductPropertiesEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RocketMqExtendProductPropertiesEntity struct{}"
	}

	return strings.Join([]string{"RocketMqExtendProductPropertiesEntity", string(data)}, " ")
}
