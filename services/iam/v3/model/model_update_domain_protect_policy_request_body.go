package model

import (
	"encoding/json"

	"strings"
)

//
type UpdateDomainProtectPolicyRequestBody struct {
	ProtectPolicy *ProtectPolicyOption `json:"protect_policy"`
}

func (o UpdateDomainProtectPolicyRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainProtectPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDomainProtectPolicyRequestBody", string(data)}, " ")
}
