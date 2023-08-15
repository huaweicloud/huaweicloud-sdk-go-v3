package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginIp 登录源IP
type LoginIp struct {
}

func (o LoginIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginIp struct{}"
	}

	return strings.Join([]string{"LoginIp", string(data)}, " ")
}
