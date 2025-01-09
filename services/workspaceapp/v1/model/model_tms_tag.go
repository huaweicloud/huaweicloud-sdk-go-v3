package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TmsTag tms标签
type TmsTag struct {

	// 键。最大长度128个unicode字符。 1. 可用 UTF-8 格式表示的字母(包含中文、西班牙语、葡语等)、数字和空格，以及以下字符： _ . : = + - @ 2. 两头不能有空白字符
	Key string `json:"key"`

	// 值。每个值最大长度255个unicode字符。 1. 可用 UTF-8 格式表示的字母(包含中文、西班牙语、葡语等)、数字和空格，以及以下字符： _ . : = + - @
	Value *string `json:"value,omitempty"`
}

func (o TmsTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsTag struct{}"
	}

	return strings.Join([]string{"TmsTag", string(data)}, " ")
}
