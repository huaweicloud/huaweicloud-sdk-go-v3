package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RabbitMqProductSupportFeaturesEntity 支持的特性功能。
type RabbitMqProductSupportFeaturesEntity struct {

	// 特性名称。
	Name *string `json:"name,omitempty"`

	// 功能特性的键值对。
	Properties map[string]string `json:"properties,omitempty"`
}

func (o RabbitMqProductSupportFeaturesEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RabbitMqProductSupportFeaturesEntity struct{}"
	}

	return strings.Join([]string{"RabbitMqProductSupportFeaturesEntity", string(data)}, " ")
}
