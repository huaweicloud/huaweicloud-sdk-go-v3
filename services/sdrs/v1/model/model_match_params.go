package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// match数据结构
type MatchParams struct {
	// 键。目前限定为resource_name，后续扩展。

	Key string `json:"key"`
	// 值。每个值最大长度255个unicode字符。

	Value string `json:"value"`
}

func (o MatchParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MatchParams struct{}"
	}

	return strings.Join([]string{"MatchParams", string(data)}, " ")
}
