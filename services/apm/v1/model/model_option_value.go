package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OptionValue 可选值
type OptionValue struct {

	// 可选值
	Value *string `json:"value,omitempty"`

	// 提示信息
	Hint *string `json:"hint,omitempty"`
}

func (o OptionValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OptionValue struct{}"
	}

	return strings.Join([]string{"OptionValue", string(data)}, " ")
}
