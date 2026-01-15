package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceInstanceResponseSysTags struct {

	// 键。最大长度128个字符。
	Key *string `json:"key,omitempty"`

	// 值。每个值最大长度255个字符。
	Value *string `json:"value,omitempty"`
}

func (o ResourceInstanceResponseSysTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceInstanceResponseSysTags struct{}"
	}

	return strings.Join([]string{"ResourceInstanceResponseSysTags", string(data)}, " ")
}
