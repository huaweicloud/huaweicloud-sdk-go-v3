package model

import (
	"encoding/json"

	"strings"
)

//
type UpdateDomainPasswordPolicyRequestBody struct {
	PasswordPolicy *PasswordPolicyOption `json:"password_policy"`
}

func (o UpdateDomainPasswordPolicyRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainPasswordPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDomainPasswordPolicyRequestBody", string(data)}, " ")
}
