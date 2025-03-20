package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePasswordPolicyV5Response Response Object
type UpdatePasswordPolicyV5Response struct {
	PasswordPolicy *PasswordPolicy `json:"password_policy,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o UpdatePasswordPolicyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePasswordPolicyV5Response struct{}"
	}

	return strings.Join([]string{"UpdatePasswordPolicyV5Response", string(data)}, " ")
}
