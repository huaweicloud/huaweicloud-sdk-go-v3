package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigItemValue 采集参数配置列表
type ConfigItemValue struct {

	// 配置项名字
	ConfigName *string `json:"config_name,omitempty"`

	// 显示名称
	DisplayName *string `json:"display_name,omitempty"`

	// 数据类型
	ConfigDataType *string `json:"config_data_type,omitempty"`

	// 最大长度
	MaxLength *int32 `json:"max_length,omitempty"`

	// 最小值
	MinValue *float64 `json:"min_value,omitempty"`

	// 最大值
	MaxValue *float64 `json:"max_value,omitempty"`

	// 可选值
	OptionValues *[]OptionValue `json:"option_values,omitempty"`

	// 默认值
	DefaultValue *string `json:"default_value,omitempty"`

	// 开始版本
	SinceVersion *string `json:"since_version,omitempty"`

	// 截至版本
	DeadVersion *string `json:"dead_version,omitempty"`

	// 值
	Value *string `json:"value,omitempty"`

	// 对象数组
	ObjectArrayPatterns *[]ObjectArrayPatterns `json:"object_array_patterns,omitempty"`

	// 实际生效值
	OverrideList *[]ConfigItemOverride `json:"override_list,omitempty"`
}

func (o ConfigItemValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigItemValue struct{}"
	}

	return strings.Join([]string{"ConfigItemValue", string(data)}, " ")
}
