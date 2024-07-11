package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RocketMqConfigResp struct {

	// RocketMQ配置名称。
	Name *string `json:"name,omitempty"`

	// RocketMQ配置当前值。
	Value *string `json:"value,omitempty"`

	// RocketMQ配置的类型。
	ConfigType *string `json:"config_type,omitempty"`

	// RocketMQ配置的默认值。
	DefaultValue *string `json:"default_value,omitempty"`

	// RocketMQ配置取值的范围。
	ValidValues *string `json:"valid_values,omitempty"`

	// RocketMQ配置值的类型。
	ValueType *string `json:"value_type,omitempty"`
}

func (o RocketMqConfigResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RocketMqConfigResp struct{}"
	}

	return strings.Join([]string{"RocketMqConfigResp", string(data)}, " ")
}
