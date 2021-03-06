package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowEnterpriseProjectQuotaResponse struct {
	Quotas         *QuotasDetail `json:"quotas,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowEnterpriseProjectQuotaResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowEnterpriseProjectQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowEnterpriseProjectQuotaResponse", string(data)}, " ")
}
