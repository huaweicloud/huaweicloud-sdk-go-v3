package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowDomainPasswordPolicyResponse struct {
	PasswordPolicy *PasswordPolicyResult `json:"password_policy,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowDomainPasswordPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDomainPasswordPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainPasswordPolicyResponse", string(data)}, " ")
}
