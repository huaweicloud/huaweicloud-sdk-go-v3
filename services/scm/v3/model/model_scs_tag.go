package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScsTag struct {

	// 标签键。 可用 UTF-8 格式表示的字母(包含中文、西班牙语、葡语等)、数字和空格，以及以下字符： _ . : = + - @
	Key *string `json:"key,omitempty"`

	// 标签值列表。 每个值可用 UTF-8 格式表示的字母(包含中文、西班牙语、葡语等)、数字和空格，以及以下字符： _ . : = + - @
	Values *[]string `json:"values,omitempty"`
}

func (o ScsTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScsTag struct{}"
	}

	return strings.Join([]string{"ScsTag", string(data)}, " ")
}
