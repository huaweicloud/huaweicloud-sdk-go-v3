package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryRedisConfig 实例配置项
type QueryRedisConfig struct {

	// 配置参数值。
	ParamValue *string `json:"param_value,omitempty"`

	// 配置参数的值类型。
	ValueType *string `json:"value_type,omitempty"`

	// 配置参数的取值范围。
	ValueRange *string `json:"value_range,omitempty"`

	// 配置项的描述。
	Description *string `json:"description,omitempty"`

	// 配置参数的默认值。
	DefaultValue *string `json:"default_value,omitempty"`

	// 配置参数名称。
	ParamName *string `json:"param_name,omitempty"`

	// 配置参数ID。
	ParamId *string `json:"param_id,omitempty"`
}

func (o QueryRedisConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryRedisConfig struct{}"
	}

	return strings.Join([]string{"QueryRedisConfig", string(data)}, " ")
}
