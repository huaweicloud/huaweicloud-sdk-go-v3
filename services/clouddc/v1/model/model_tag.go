package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tag 标签对象
type Tag struct {

	// 键。最大长度128个字符。 key需要满足3.1 KEY字符集规范。
	Key string `json:"key"`

	// 值,每个值最大长度255个字符。
	Value *string `json:"value,omitempty"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
