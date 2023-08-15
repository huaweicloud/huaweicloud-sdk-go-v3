package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomVariable struct {

	// 流水线ID
	PipelineId *string `json:"pipeline_id,omitempty"`

	// 自定义参数名称
	Name *string `json:"name,omitempty"`

	// 自定义参数顺序
	Sequence *int32 `json:"sequence,omitempty"`

	// 自定义参数类型
	Type *string `json:"type,omitempty"`

	// 自定义参数默认值
	Value *string `json:"value,omitempty"`

	// 是否私密参数
	IsSecret *bool `json:"is_secret,omitempty"`

	// 自定义参数描述
	Description *string `json:"description,omitempty"`

	// 是否运行时设置
	IsRuntime *bool `json:"is_runtime,omitempty"`

	// 参数限制
	Limits *[]interface{} `json:"limits,omitempty"`

	// 是否重置
	IsReset *bool `json:"is_reset,omitempty"`

	// 最近一次参数值
	LatestValue *string `json:"latest_value,omitempty"`

	// 运行时传入值
	RuntimeValue *string `json:"runtime_value,omitempty"`
}

func (o CustomVariable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomVariable struct{}"
	}

	return strings.Join([]string{"CustomVariable", string(data)}, " ")
}
