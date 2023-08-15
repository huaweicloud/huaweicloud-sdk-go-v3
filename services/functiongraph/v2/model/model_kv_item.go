package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KvItem 键值对
type KvItem struct {

	// 键
	Key *string `json:"key,omitempty"`

	// 值
	Value *string `json:"value,omitempty"`
}

func (o KvItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KvItem struct{}"
	}

	return strings.Join([]string{"KvItem", string(data)}, " ")
}
