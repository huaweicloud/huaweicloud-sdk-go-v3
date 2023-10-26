package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineByTemplateDtoVariables struct {

	// 参数名称
	Name *string `json:"name,omitempty"`

	// 参数序号
	Sequence *int32 `json:"sequence,omitempty"`

	// 参数类型
	Type *string `json:"type,omitempty"`

	// 参数值
	Value *string `json:"value,omitempty"`

	// 是否私密参数
	IsSecret *bool `json:"is_secret,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 是否运行时设置
	IsRuntime *bool `json:"is_runtime,omitempty"`

	// 是否重置
	IsReset *bool `json:"is_reset,omitempty"`

	// 最后一次参数值
	LatestValue *string `json:"latest_value,omitempty"`

	// 枚举值
	Limits *[]string `json:"limits,omitempty"`
}

func (o PipelineByTemplateDtoVariables) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineByTemplateDtoVariables struct{}"
	}

	return strings.Join([]string{"PipelineByTemplateDtoVariables", string(data)}, " ")
}
