package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectResourceTagResponseTags struct {

	// 键。最大长度128个字符。 key满足3.1 KEY字符集规范。
	Key string `json:"key"`

	// 值列表。每个值最大长度255个字符。 value满足3.2 VALUE字符集规范。
	Values []string `json:"values"`
}

func (o ProjectResourceTagResponseTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectResourceTagResponseTags struct{}"
	}

	return strings.Join([]string{"ProjectResourceTagResponseTags", string(data)}, " ")
}
