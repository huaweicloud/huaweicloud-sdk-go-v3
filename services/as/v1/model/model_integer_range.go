package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IntegerRange struct {

	// 伸缩组最大实例数
	Max *int32 `json:"max,omitempty"`

	// 伸缩组最小实例数
	Min *int32 `json:"min,omitempty"`

	// 伸缩组期望实例数
	Desire *int32 `json:"desire,omitempty"`
}

func (o IntegerRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IntegerRange struct{}"
	}

	return strings.Join([]string{"IntegerRange", string(data)}, " ")
}
