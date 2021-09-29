package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowGaussMySqlProjectQuotasResponse struct {
	Quotas         *ProjectQuotas `json:"quotas,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowGaussMySqlProjectQuotasResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlProjectQuotasResponse struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlProjectQuotasResponse", string(data)}, " ")
}
