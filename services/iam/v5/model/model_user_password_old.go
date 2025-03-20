package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserPasswordOld IAM用户的旧密码。
type UserPasswordOld struct {
}

func (o UserPasswordOld) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserPasswordOld struct{}"
	}

	return strings.Join([]string{"UserPasswordOld", string(data)}, " ")
}
