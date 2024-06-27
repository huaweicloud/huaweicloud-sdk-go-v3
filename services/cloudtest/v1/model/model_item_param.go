package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ItemParam struct {
	BasicValue *AwParamBasicValue `json:"basic_value,omitempty"`

	// 默认值
	DefaultValue *string `json:"defaultValue,omitempty"`

	// 枚举类型
	EnumType *string `json:"enum_type,omitempty"`

	// 当前选中的枚举类型
	EnumTypeSelected *string `json:"enum_type_selected,omitempty"`

	// num和String有效
	Format *string `json:"format,omitempty"`

	// 是否是大字段参数
	IsBigField *bool `json:"is_big_field,omitempty"`

	// Array类型最大元素数
	MaxItems *int32 `json:"maxItems,omitempty"`

	// String类型最大长度
	MaxLength *int64 `json:"maxLength,omitempty"`

	Maximum *Number `json:"maximum,omitempty"`

	// Array类型最小元素数
	MinItems *int32 `json:"minItems,omitempty"`

	// String类型最小长度
	MinLength *int64 `json:"minLength,omitempty"`

	Minimum *Number `json:"minimum,omitempty"`

	// num和String有效
	Pattern *string `json:"pattern,omitempty"`

	// 参数类型
	Type *string `json:"type,omitempty"`

	ValidateRule *ValidateRule `json:"validate_rule,omitempty"`

	// 参数值类型 0-基本类型，type字段为String,Integer等基本类型 1-基本类型数组,type字段为List<String>,List<Integer>等基本类型List 2-结构体，type字段为除了基本类型以外的结构体 3-结构体数组，type字段为List<结构体> 5-前端枚举类型
	ValueType *int32 `json:"value_type,omitempty"`

	// choice选择关系
	XChoiceValue *string `json:"xChoiceValue,omitempty"`
}

func (o ItemParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ItemParam struct{}"
	}

	return strings.Join([]string{"ItemParam", string(data)}, " ")
}
