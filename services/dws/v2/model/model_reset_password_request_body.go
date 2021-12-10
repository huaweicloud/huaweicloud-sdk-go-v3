package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type ResetPasswordRequestBody struct {
	// DWS集群管理员新密码

	NewPassword string `json:"new_password"`
}

func (o ResetPasswordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPasswordRequestBody struct{}"
	}

	return strings.Join([]string{"ResetPasswordRequestBody", string(data)}, " ")
}
