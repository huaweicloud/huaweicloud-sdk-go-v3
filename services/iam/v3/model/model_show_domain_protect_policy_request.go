package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDomainProtectPolicyRequest struct {
	DomainId string `json:"domain_id"`
}

func (o ShowDomainProtectPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDomainProtectPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainProtectPolicyRequest", string(data)}, " ")
}
