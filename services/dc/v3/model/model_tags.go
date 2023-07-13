package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Tags struct {

	// 键。最大长度127个unicode字符。 key不能为空。
	Key string `json:"key"`

	// 值列表。每个值最大长度255个unicode字符。
	Values []string `json:"values"`
}

func (o Tags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tags struct{}"
	}

	return strings.Join([]string{"Tags", string(data)}, " ")
}
