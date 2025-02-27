package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeUsersPassword struct {

	// 新密码。
	Password string `json:"password"`

	// 旧密码。
	OldPassword string `json:"old_password"`
}

func (o ChangeUsersPassword) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeUsersPassword struct{}"
	}

	return strings.Join([]string{"ChangeUsersPassword", string(data)}, " ")
}
