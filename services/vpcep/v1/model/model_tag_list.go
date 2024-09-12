package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagList 标签列表，没有标签默认为空数组。
type TagList struct {

	// 键。 key不能为空，长度1~128个字符（中文也可以输入128个字符）。  可用 UTF-8 格式表示的字母(包含中文、西班牙语、葡语等)、数字和空格，以及以下字符： _ . : = + - @。 _sys_开头属于系统标签，租户不能输入。 key两头不能有空格字符。
	Key *string `json:"key,omitempty"`

	// 值。 长度0~255个字符（中文也可以输入255个字符）。  可用 UTF-8 格式表示的字母(包含中文、西班牙语、葡语等)、数字和空格，以及以下字符： _ . : / = + - @。 资源标签值可以为空字符串。
	Value *string `json:"value,omitempty"`
}

func (o TagList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagList struct{}"
	}

	return strings.Join([]string{"TagList", string(data)}, " ")
}
