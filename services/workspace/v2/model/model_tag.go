package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tag 标签。
type Tag struct {

	// 标签的键，不能为空，最大长度128个unicode字符。标签的键可以包含任意语种字母、数字、空格和_.:=+-@，但首尾不能含有空格，不能以_sys_开头。
	Key string `json:"key"`

	// 标签的值，最大长度43个unicode字符。标签的值可以包含任意语种字母、数字、空格和_.:=+-@，但首尾不能含有空格。
	Value *string `json:"value,omitempty"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
