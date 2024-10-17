package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceTagDeleteRequestTags struct {

	// 键。最大长度128个字符。
	Key string `json:"key"`

	// 值。每个值最大长度255个字符。
	Value *string `json:"value,omitempty"`
}

func (o ResourceTagDeleteRequestTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTagDeleteRequestTags struct{}"
	}

	return strings.Join([]string{"ResourceTagDeleteRequestTags", string(data)}, " ")
}
