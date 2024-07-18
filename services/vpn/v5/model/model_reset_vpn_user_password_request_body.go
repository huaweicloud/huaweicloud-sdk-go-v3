package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResetVpnUserPasswordRequestBody struct {

	// 用户新密码
	NewPassword string `json:"new_password"`
}

func (o ResetVpnUserPasswordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetVpnUserPasswordRequestBody struct{}"
	}

	return strings.Join([]string{"ResetVpnUserPasswordRequestBody", string(data)}, " ")
}
