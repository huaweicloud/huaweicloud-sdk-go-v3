package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVpnUserPasswordRequestBody struct {

	// 用户旧密码
	OldPassword string `json:"old_password"`

	// 用户新密码
	NewPassword string `json:"new_password"`
}

func (o UpdateVpnUserPasswordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnUserPasswordRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVpnUserPasswordRequestBody", string(data)}, " ")
}
