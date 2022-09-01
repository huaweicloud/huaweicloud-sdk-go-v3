package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 聚合计算
type AggregateResponse struct {

	// 输入参数
	Inputs *[]InputResponse `json:"inputs,omitempty" xml:"inputs"`

	Expression *Expression `json:"expression,omitempty" xml:"expression"`

	// 输出属性名(不推荐使用，待废弃，使用outputs替代)
	OutputProperty *string `json:"output_property,omitempty" xml:"output_property"`

	Outputs *[]OutputResponse `json:"outputs,omitempty" xml:"outputs"`

	Schedule *DtSchedule `json:"schedule,omitempty" xml:"schedule"`
}

func (o AggregateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregateResponse struct{}"
	}

	return strings.Join([]string{"AggregateResponse", string(data)}, " ")
}
