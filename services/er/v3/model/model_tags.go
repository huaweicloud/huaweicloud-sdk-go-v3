package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签
type Tags struct {

	// 标签键，不能为空，最大长度127个unicode字符。
	Key string `json:"key"`

	// 值列表，每个值最大长度255个unicode字符。
	Values []string `json:"values"`
}

func (o Tags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tags struct{}"
	}

	return strings.Join([]string{"Tags", string(data)}, " ")
}
