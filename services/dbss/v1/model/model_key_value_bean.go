package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KeyValueBean struct {

	// 键
	Key string `json:"key"`

	// 值
	Value string `json:"value"`
}

func (o KeyValueBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeyValueBean struct{}"
	}

	return strings.Join([]string{"KeyValueBean", string(data)}, " ")
}
