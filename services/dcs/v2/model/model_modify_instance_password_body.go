package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyInstancePasswordBody struct {
	// 旧密码

	OldPassword *string `json:"old_password,omitempty"`
	// 新密码

	NewPassword *string `json:"new_password,omitempty"`
}

func (o ModifyInstancePasswordBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyInstancePasswordBody struct{}"
	}

	return strings.Join([]string{"ModifyInstancePasswordBody", string(data)}, " ")
}
