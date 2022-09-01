package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 转换计算
type TransformResponse struct {

	// 输入参数
	Inputs *[]InputResponse `json:"inputs,omitempty" xml:"inputs"`

	Expression *Formula `json:"expression,omitempty" xml:"expression"`

	// 输出属性名(不推荐使用，待废弃，使用outputs替代)
	OutputProperty *string `json:"output_property,omitempty" xml:"output_property"`

	Outputs *[]OutputResponse `json:"outputs,omitempty" xml:"outputs"`
}

func (o TransformResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransformResponse struct{}"
	}

	return strings.Join([]string{"TransformResponse", string(data)}, " ")
}
