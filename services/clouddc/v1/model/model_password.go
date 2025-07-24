package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Password 设置实例的管理员账户初始登录密码，其中，Linux管理员账户为root，Windows管理员账户为Administrator。
type Password struct {
}

func (o Password) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Password struct{}"
	}

	return strings.Join([]string{"Password", string(data)}, " ")
}
