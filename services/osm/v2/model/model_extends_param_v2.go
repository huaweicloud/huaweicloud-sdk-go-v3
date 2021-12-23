package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtendsParamV2 struct {
	// 提示

	Tips *string `json:"tips,omitempty"`
	// 是否必填

	Required *int32 `json:"required,omitempty"`
	// 限制长度

	Length *int32 `json:"length,omitempty"`
	// 语言

	Language *string `json:"language,omitempty"`
	// 参数标识

	ParamKey *string `json:"param_key,omitempty"`
	// 参数名称

	ParamName *string `json:"param_name,omitempty"`
	// 参数类型

	ParamType *int32 `json:"param_type,omitempty"`
	// 参数描述

	ParamDesc *string `json:"param_desc,omitempty"`
	// 默认值

	DefaultValue *string `json:"default_value,omitempty"`
	// 最大值

	MaxValue *int64 `json:"max_value,omitempty"`
	// 最小值

	MinValue *int64 `json:"min_value,omitempty"`
	// 选项值

	SelectItem *string `json:"select_item,omitempty"`
	// 是否展示

	IsShow *int32 `json:"is_show,omitempty"`
	// 是否敏感

	IsSensitive *int32 `json:"is_sensitive,omitempty"`
}

func (o ExtendsParamV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendsParamV2 struct{}"
	}

	return strings.Join([]string{"ExtendsParamV2", string(data)}, " ")
}
