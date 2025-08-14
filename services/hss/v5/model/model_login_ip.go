package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginIp **参数解释**： 登录源IP **取值范围**： 字符长度1-256位
type LoginIp struct {
}

func (o LoginIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginIp struct{}"
	}

	return strings.Join([]string{"LoginIp", string(data)}, " ")
}
