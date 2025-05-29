package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScsMatch struct {

	// 标签键。 可用 UTF-8 格式表示的字母(包含中文、西班牙语、葡语等)、数字和空格，以及以下字符： _ . : = + - @
	Key *string `json:"key,omitempty"`

	// 标签值。 可用 UTF-8 格式表示的字母(包含中文、西班牙语、葡语等)、数字和空格，以及以下字符： _ . : / = + - @
	Value *string `json:"value,omitempty"`
}

func (o ScsMatch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScsMatch struct{}"
	}

	return strings.Join([]string{"ScsMatch", string(data)}, " ")
}
