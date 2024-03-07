package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PathElement 策略的JSON表示的路径的单个元素。
type PathElement struct {

	// 数组中的索引，从0开始。
	Index *int32 `json:"index,omitempty"`

	// 对象中的键。
	Key *string `json:"key,omitempty"`

	Substring *Substring `json:"substring,omitempty"`

	// 与对象中给定键关联的值。
	Value *string `json:"value,omitempty"`
}

func (o PathElement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PathElement struct{}"
	}

	return strings.Join([]string{"PathElement", string(data)}, " ")
}
