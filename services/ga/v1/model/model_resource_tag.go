package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签。
type ResourceTag struct {

	// 标签键，规则： - 1.key不能为空，长度1~128个字符（中文也可以输入128个字符）。 - 2.可用 UTF-8 格式表示的字母（包含中文）、数字和空格，以及以下字符_ . : = + - @。 - 3._sys_开头属于系统标签，租户不能输入。
	Key string `json:"key"`

	// 标签值，规则： - 1.长度0~255个字符（中文也可以输入255个字符）。 - 2.可用 UTF-8 格式表示的字母（包含中文）、数字和空格，以及以下字符_ . : / = + - @。 - 3.资源标签值可以为空（empty or null）。
	Value string `json:"value"`
}

func (o ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
