package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPasswordPolicyV5Response Response Object
type ShowPasswordPolicyV5Response struct {
	PasswordPolicy *PasswordPolicy `json:"password_policy,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowPasswordPolicyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPasswordPolicyV5Response struct{}"
	}

	return strings.Join([]string{"ShowPasswordPolicyV5Response", string(data)}, " ")
}
