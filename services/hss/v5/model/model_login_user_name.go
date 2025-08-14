package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginUserName **参数解释**： 登录用户名 **取值范围**： 字符长度1-256位
type LoginUserName struct {
}

func (o LoginUserName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginUserName struct{}"
	}

	return strings.Join([]string{"LoginUserName", string(data)}, " ")
}
