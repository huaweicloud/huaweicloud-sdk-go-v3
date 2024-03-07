package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Substring JSON反序列化后的字符串的子串。
type Substring struct {

	// 子字符串的起始索引，从0开始。0表示第一个字符。
	Start int32 `json:"start"`

	// 子字符串的长度。
	Length int32 `json:"length"`
}

func (o Substring) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Substring struct{}"
	}

	return strings.Join([]string{"Substring", string(data)}, " ")
}
