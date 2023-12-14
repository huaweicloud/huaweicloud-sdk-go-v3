package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineVariable 流水线自定义参数
type PipelineVariable struct {

	// 参数名称
	Name *string `json:"name,omitempty"`

	// 序号
	Sequence *int32 `json:"sequence,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 默认值
	Value *string `json:"value,omitempty"`

	// 是否私密
	IsSecret *bool `json:"is_secret,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 是否运行时设置
	IsRuntime *bool `json:"is_runtime,omitempty"`

	// 限定枚举值
	Limits *[]string `json:"limits,omitempty"`

	// 自增长参数是否被重置
	IsReset *bool `json:"is_reset,omitempty"`

	// 自增长参数最新值
	LatestValue *string `json:"latest_value,omitempty"`
}

func (o PipelineVariable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineVariable struct{}"
	}

	return strings.Join([]string{"PipelineVariable", string(data)}, " ")
}
