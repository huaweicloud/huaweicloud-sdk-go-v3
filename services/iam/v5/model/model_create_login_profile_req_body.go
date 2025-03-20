package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLoginProfileReqBody The request body of create login profile request.
type CreateLoginProfileReqBody struct {

	// IAM用户的密码。
	Password string `json:"password"`

	// IAM用户下次登录时是否需要修改密码。
	PasswordResetRequired bool `json:"password_reset_required"`
}

func (o CreateLoginProfileReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoginProfileReqBody struct{}"
	}

	return strings.Join([]string{"CreateLoginProfileReqBody", string(data)}, " ")
}
