package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateDomainProtectPolicyRequest struct {
	DomainId string                                `json:"domain_id"`
	Body     *UpdateDomainProtectPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateDomainProtectPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainProtectPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainProtectPolicyRequest", string(data)}, " ")
}
