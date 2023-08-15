package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValueRange 数值值域区间定义
type ValueRange struct {

	// 值域下限
	Lower float32 `json:"lower,omitempty"`

	// 是否包含值域下限
	LowerInclusive *bool `json:"lower_inclusive,omitempty"`

	// 值域上限
	Upper float32 `json:"upper,omitempty"`

	// 是否包含值域上限
	UpperInclusive *bool `json:"upper_inclusive,omitempty"`
}

func (o ValueRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueRange struct{}"
	}

	return strings.Join([]string{"ValueRange", string(data)}, " ")
}
