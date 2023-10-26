package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LoginWebCliBody struct {

	// 登录Redis密码
	Password *string `json:"password,omitempty"`
}

func (o LoginWebCliBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginWebCliBody struct{}"
	}

	return strings.Join([]string{"LoginWebCliBody", string(data)}, " ")
}
