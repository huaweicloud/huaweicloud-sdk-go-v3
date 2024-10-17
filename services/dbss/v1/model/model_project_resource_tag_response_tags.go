package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectResourceTagResponseTags struct {

	// 键。最大长度128个字符。
	Key *string `json:"key,omitempty"`

	// 值列表。每个值最大长度255个字符。
	Values *[]string `json:"values,omitempty"`
}

func (o ProjectResourceTagResponseTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectResourceTagResponseTags struct{}"
	}

	return strings.Join([]string{"ProjectResourceTagResponseTags", string(data)}, " ")
}
