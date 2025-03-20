package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangePasswordReqBody The request body of change password request.
type ChangePasswordReqBody struct {

	// IAM用户的新密码。
	NewPassword string `json:"new_password"`

	// IAM用户的旧密码。
	OldPassword string `json:"old_password"`
}

func (o ChangePasswordReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePasswordReqBody struct{}"
	}

	return strings.Join([]string{"ChangePasswordReqBody", string(data)}, " ")
}
