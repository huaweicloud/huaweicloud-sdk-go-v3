package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PasswordResetRequired IAM用户下次登录时是否需要修改密码。
type PasswordResetRequired struct {
}

func (o PasswordResetRequired) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PasswordResetRequired struct{}"
	}

	return strings.Join([]string{"PasswordResetRequired", string(data)}, " ")
}
