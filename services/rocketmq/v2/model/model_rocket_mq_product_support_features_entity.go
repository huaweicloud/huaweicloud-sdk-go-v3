package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RocketMqProductSupportFeaturesEntity 支持的特性功能。
type RocketMqProductSupportFeaturesEntity struct {

	// 特性名称。
	Name *string `json:"name,omitempty"`

	// 功能特性的键值对。
	Properties map[string]string `json:"properties,omitempty"`
}

func (o RocketMqProductSupportFeaturesEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RocketMqProductSupportFeaturesEntity struct{}"
	}

	return strings.Join([]string{"RocketMqProductSupportFeaturesEntity", string(data)}, " ")
}
