package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AwParam struct {
	BasicValue *AwParamBasicValue `json:"basic_value,omitempty"`

	// valueType为1时该值有效
	BasicValueList *[]AwParamBasicValue `json:"basic_value_list,omitempty"`

	// 自定义requestBody内容,inType为2且isBodyCustom为1时有效
	CustomBody *string `json:"custom_body,omitempty"`

	// 默认值
	DefaultValue *string `json:"defaultValue,omitempty"`

	// aw参数描述
	Description *string `json:"description,omitempty"`

	// 是否禁用 只有非必需参数才能被禁用
	Disabled *bool `json:"disabled,omitempty"`

	// 用于存储下拉菜单的值
	DropDownValue *string `json:"drop_down_value,omitempty"`

	// 枚举类型
	EnumType *string `json:"enum_type,omitempty"`

	// 当前选中的枚举类型
	EnumTypeSelected *string `json:"enum_type_selected,omitempty"`

	// num和String有效
	Format *string `json:"format,omitempty"`

	// rest接口输入参数类型 0-query 1-path 2-body 3-header 4-formdata 5-config
	InType *int32 `json:"in_type,omitempty"`

	// 是否是被选中参数
	IsChecked *bool `json:"isChecked,omitempty"`

	// 是否是大字段参数
	IsBigField *bool `json:"is_big_field,omitempty"`

	// requestBody是否自定义，inType为2时有效。0-非自定义，用原有逻辑；1-自定义body,requestBody直接使用字段customBody
	IsBodyCustom *int32 `json:"is_body_custom,omitempty"`

	// 是否是契约AW 0-否；1-yes
	IsContractParam *int32 `json:"is_contract_param,omitempty"`

	// 是否敏感参数，0  是，  1 不是
	IsSensitive *bool `json:"is_sensitive,omitempty"`

	Item *ItemParam `json:"item,omitempty"`

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

	Mock *MockInfo `json:"mock,omitempty"`

	// 参数名
	Name *string `json:"name,omitempty"`

	// valueType为2时该值有效。valueType为3时，该值用来表示数组中单个结构体
	ObjectValue *[]AwParam `json:"object_value,omitempty"`

	// valueType为3时该值有效
	ObjectValueList *[][]AwParam `json:"object_value_list,omitempty"`

	// num和String有效
	Pattern *string `json:"pattern,omitempty"`

	// 是否必需参数
	Required *bool `json:"required,omitempty"`

	// 参数类型
	Type *string `json:"type,omitempty"`

	ValidateRule *ValidateRule `json:"validate_rule,omitempty"`

	// 参数值类型 0-基本类型，type字段为String,Integer等基本类型 1-基本类型数组,type字段为List<String>,List<Integer>等基本类型List 2-结构体，type字段为除了基本类型以外的结构体 3-结构体数组，type字段为List<结构体> 5-前端枚举类型
	ValueType *int32 `json:"value_type,omitempty"`

	// choice选择关系
	XChoiceValue *string `json:"xChoiceValue,omitempty"`
}

func (o AwParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AwParam struct{}"
	}

	return strings.Join([]string{"AwParam", string(data)}, " ")
}
