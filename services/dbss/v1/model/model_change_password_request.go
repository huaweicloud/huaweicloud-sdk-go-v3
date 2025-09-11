package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangePasswordRequest struct {

	// 新密码
	NewPassword string `json:"new_password"`

	// 用户名
	UserName string `json:"user_name"`
}

func (o ChangePasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePasswordRequest struct{}"
	}

	return strings.Join([]string{"ChangePasswordRequest", string(data)}, " ")
}
