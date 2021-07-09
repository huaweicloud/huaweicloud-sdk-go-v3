package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowMysqlProjectQuotasResponse struct {
	Quotas         *ProjectQuotas `json:"quotas,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowMysqlProjectQuotasResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMysqlProjectQuotasResponse struct{}"
	}

	return strings.Join([]string{"ShowMysqlProjectQuotasResponse", string(data)}, " ")
}
