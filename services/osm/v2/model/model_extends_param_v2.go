package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtendsParamV2 struct {

	// 提示
	Tips *string `json:"tips,omitempty" xml:"tips"`

	// 是否必填
	Required *int32 `json:"required,omitempty" xml:"required"`

	// 限制长度
	Length *int32 `json:"length,omitempty" xml:"length"`

	// 语言
	Language *string `json:"language,omitempty" xml:"language"`

	// 参数标识
	ParamKey *string `json:"param_key,omitempty" xml:"param_key"`

	// 参数名称
	ParamName *string `json:"param_name,omitempty" xml:"param_name"`

	// 参数类型
	ParamType *int32 `json:"param_type,omitempty" xml:"param_type"`

	// 参数描述
	ParamDesc *string `json:"param_desc,omitempty" xml:"param_desc"`

	// 默认值
	DefaultValue *string `json:"default_value,omitempty" xml:"default_value"`

	// 最大值
	MaxValue *int64 `json:"max_value,omitempty" xml:"max_value"`

	// 最小值
	MinValue *int64 `json:"min_value,omitempty" xml:"min_value"`

	// 选项值
	SelectItem *string `json:"select_item,omitempty" xml:"select_item"`

	// 是否展示
	IsShow *int32 `json:"is_show,omitempty" xml:"is_show"`

	// 是否敏感
	IsSensitive *int32 `json:"is_sensitive,omitempty" xml:"is_sensitive"`
}

func (o ExtendsParamV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendsParamV2 struct{}"
	}

	return strings.Join([]string{"ExtendsParamV2", string(data)}, " ")
}
