package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicMatch Match字段数据结构说明
type PublicMatch struct {

	// 键。限定为resource_name。
	Key string `json:"key"`

	// 值。每个值最大长度255个unicode字符。
	Value string `json:"value"`
}

func (o PublicMatch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicMatch struct{}"
	}

	return strings.Join([]string{"PublicMatch", string(data)}, " ")
}
