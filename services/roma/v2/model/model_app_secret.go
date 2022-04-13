package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 应用认证访问SECRET,未提供（字段不存在或值为null）时随机生成 - 字符集：英文字母、数字、！、@、#、$、%、+、=、点、中划线、斜线/ - 复杂度：满足大小写字母、数字、特殊字符的复杂度组合，考虑兼容性暂时可不做
type AppSecret struct {
}

func (o AppSecret) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppSecret struct{}"
	}

	return strings.Join([]string{"AppSecret", string(data)}, " ")
}
