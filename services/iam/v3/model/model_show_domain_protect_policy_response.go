package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowDomainProtectPolicyResponse struct {
	ProtectPolicy  *ProtectPolicyResult `json:"protect_policy,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowDomainProtectPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDomainProtectPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainProtectPolicyResponse", string(data)}, " ")
}
