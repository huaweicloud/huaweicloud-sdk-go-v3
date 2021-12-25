package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 应用认证访问KEY,未提供时随机生成 - 字符集：支持中文、英文字母、数字、中划线、下划线、@号和点，以字母或中文或数字开头 - 约束：实例下唯一
type AppKey struct {
}

func (o AppKey) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppKey struct{}"
	}

	return strings.Join([]string{"AppKey", string(data)}, " ")
}
