package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlgorithmDetailDto 动态脱敏策略算法详情
type AlgorithmDetailDto struct {

	// 开始位数，该值需要大于0且小于end值。
	Start *int32 `json:"start,omitempty"`

	// 结束位数，该值需要大于或等于start值。
	End *int32 `json:"end,omitempty"`

	// 数值类型。
	IntTarget *int32 `json:"int_target,omitempty"`

	// 字符串类型。可输入参数为\"*\"或者\"#\"。
	StringTarget *string `json:"string_target,omitempty"`
}

func (o AlgorithmDetailDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlgorithmDetailDto struct{}"
	}

	return strings.Join([]string{"AlgorithmDetailDto", string(data)}, " ")
}
