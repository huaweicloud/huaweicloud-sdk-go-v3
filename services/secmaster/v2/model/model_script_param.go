package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScriptParam struct {

	// 键
	Key *string `json:"key,omitempty"`

	// 值
	Value *string `json:"value,omitempty"`
}

func (o ScriptParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptParam struct{}"
	}

	return strings.Join([]string{"ScriptParam", string(data)}, " ")
}
