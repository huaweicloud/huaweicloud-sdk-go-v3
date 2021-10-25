package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowDomainQuotaResponse struct {
	Quotas         *[]DomainQuotaResponseQuotas `json:"quotas,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowDomainQuotaResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDomainQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainQuotaResponse", string(data)}, " ")
}
