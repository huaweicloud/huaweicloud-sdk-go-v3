package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLoginProfileReqBody The request body of update login profile request.
type UpdateLoginProfileReqBody struct {

	// IAM用户的密码。
	Password *string `json:"password,omitempty"`

	// IAM用户下次登录时是否需要修改密码。
	PasswordResetRequired *bool `json:"password_reset_required,omitempty"`
}

func (o UpdateLoginProfileReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoginProfileReqBody struct{}"
	}

	return strings.Join([]string{"UpdateLoginProfileReqBody", string(data)}, " ")
}
