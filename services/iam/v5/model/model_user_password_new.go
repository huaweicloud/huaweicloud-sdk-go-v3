package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserPasswordNew IAM用户的新密码。
type UserPasswordNew struct {
}

func (o UserPasswordNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserPasswordNew struct{}"
	}

	return strings.Join([]string{"UserPasswordNew", string(data)}, " ")
}
