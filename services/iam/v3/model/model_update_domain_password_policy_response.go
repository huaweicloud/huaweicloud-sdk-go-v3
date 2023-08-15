package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainPasswordPolicyResponse Response Object
type UpdateDomainPasswordPolicyResponse struct {
	PasswordPolicy *PasswordPolicyResult `json:"password_policy,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o UpdateDomainPasswordPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainPasswordPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainPasswordPolicyResponse", string(data)}, " ")
}
