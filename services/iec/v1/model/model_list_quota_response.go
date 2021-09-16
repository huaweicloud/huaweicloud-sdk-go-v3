package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListQuotaResponse struct {
	Quotas         *QuotaResources `json:"quotas,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListQuotaResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListQuotaResponse struct{}"
	}

	return strings.Join([]string{"ListQuotaResponse", string(data)}, " ")
}
