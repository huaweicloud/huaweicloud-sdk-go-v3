package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletingResourceTag 标签。
type DeletingResourceTag struct {

	// 标签键，取值范围： - key不能为空，长度1~128个字符（中文也可以输入128个字符）。 - 可用 UTF-8 格式表示的字母（包含中文）、数字和空格，以及以下字符_ . : = + - @。 - _sys_开头属于系统标签，租户不能输入。
	Key string `json:"key"`
}

func (o DeletingResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletingResourceTag struct{}"
	}

	return strings.Join([]string{"DeletingResourceTag", string(data)}, " ")
}
