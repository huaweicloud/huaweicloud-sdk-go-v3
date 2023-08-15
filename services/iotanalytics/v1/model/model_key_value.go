package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KeyValue struct {

	// 键。
	Key *string `json:"key,omitempty"`

	// 值。
	Value *string `json:"value,omitempty"`
}

func (o KeyValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeyValue struct{}"
	}

	return strings.Join([]string{"KeyValue", string(data)}, " ")
}
