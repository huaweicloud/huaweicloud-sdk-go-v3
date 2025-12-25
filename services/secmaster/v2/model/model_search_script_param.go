package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SearchScriptParam struct {

	// 键
	Key *string `json:"key,omitempty"`

	// 值
	Value *string `json:"value,omitempty"`
}

func (o SearchScriptParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchScriptParam struct{}"
	}

	return strings.Join([]string{"SearchScriptParam", string(data)}, " ")
}
