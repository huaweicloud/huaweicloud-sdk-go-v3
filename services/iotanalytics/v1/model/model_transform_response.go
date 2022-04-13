package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 转换计算
type TransformResponse struct {
	// 输入参数

	Inputs *[]InputResponse `json:"inputs,omitempty"`

	Expression *Formula `json:"expression,omitempty"`
	// 输出属性名(不推荐使用，待废弃，使用outputs替代)

	OutputProperty *string `json:"output_property,omitempty"`

	Outputs *[]OutputResponse `json:"outputs,omitempty"`
}

func (o TransformResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransformResponse struct{}"
	}

	return strings.Join([]string{"TransformResponse", string(data)}, " ")
}
