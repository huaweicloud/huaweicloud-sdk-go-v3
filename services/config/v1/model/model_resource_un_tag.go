package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceUnTag struct {

	// 键。最大长度128个unicode字符。 key符合3.1 KEY字符集规范。
	Key string `json:"key"`

	// 值。每个值最大长度255个unicode字符。value符合3.2 VALUE字符集规范。
	Value string `json:"value"`
}

func (o ResourceUnTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceUnTag struct{}"
	}

	return strings.Join([]string{"ResourceUnTag", string(data)}, " ")
}
