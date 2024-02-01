package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RequiredTag 一个key/value键值对
type RequiredTag struct {

	// 键
	Key string `json:"key"`

	// 值
	Value string `json:"value"`
}

func (o RequiredTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequiredTag struct{}"
	}

	return strings.Join([]string{"RequiredTag", string(data)}, " ")
}
