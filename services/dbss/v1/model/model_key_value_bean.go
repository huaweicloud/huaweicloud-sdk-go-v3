package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KeyValueBean struct {

	// 键
	Key string `json:"key"`

	// 值
	Value *string `json:"value,omitempty"`
}

func (o KeyValueBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeyValueBean struct{}"
	}

	return strings.Join([]string{"KeyValueBean", string(data)}, " ")
}
