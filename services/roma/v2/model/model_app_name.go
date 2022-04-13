package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 应用名称 - 字符集：支持中文、英文字母、数字、中划线、下划线、点、空格和中英文圆括号 - 约束：实例下唯一
type AppName struct {
}

func (o AppName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppName struct{}"
	}

	return strings.Join([]string{"AppName", string(data)}, " ")
}
