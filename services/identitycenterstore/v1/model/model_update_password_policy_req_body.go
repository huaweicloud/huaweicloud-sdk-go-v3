package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePasswordPolicyReqBody struct {
	PasswordPolicy *PasswordPolicyDto `json:"password_policy"`
}

func (o UpdatePasswordPolicyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePasswordPolicyReqBody struct{}"
	}

	return strings.Join([]string{"UpdatePasswordPolicyReqBody", string(data)}, " ")
}
