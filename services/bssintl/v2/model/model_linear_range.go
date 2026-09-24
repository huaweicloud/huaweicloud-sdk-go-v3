package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LinearRange struct {

	// 线性属性的最小值
	Min *int32 `json:"min,omitempty"`

	// 线性属性的最大值
	Max *int32 `json:"max,omitempty"`

	// 线性属性的步长
	Step *int32 `json:"step,omitempty"`
}

func (o LinearRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinearRange struct{}"
	}

	return strings.Join([]string{"LinearRange", string(data)}, " ")
}
