package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowPolicyAndInstanceQuotaResponse struct {
	Quotas         *PolicyInstanceQuotas `json:"quotas,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowPolicyAndInstanceQuotaResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPolicyAndInstanceQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowPolicyAndInstanceQuotaResponse", string(data)}, " ")
}
