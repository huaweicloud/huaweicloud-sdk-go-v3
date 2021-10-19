package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListQuotaDetailsResponse struct {
	Quotas         *ResourcesResp `json:"quotas,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListQuotaDetailsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListQuotaDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListQuotaDetailsResponse", string(data)}, " ")
}
