package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据库用户名密码信息。
type ResetDatabasePassword struct {

	// 数据库用户名。
	Name string `json:"name"`

	// 主机地址。
	Host string `json:"host"`

	// 数据库用户密码，非空，至少包含以下字符中的三种：大写字母、小写字母、数字和特殊符号~!@#$%^*-_=+?,()&组成，长度8~32个字符。建议您输入高强度密码，以提高安全性，防止出现密码被暴力破解等安全风险
	Password string `json:"password"`
}

func (o ResetDatabasePassword) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetDatabasePassword struct{}"
	}

	return strings.Join([]string{"ResetDatabasePassword", string(data)}, " ")
}
