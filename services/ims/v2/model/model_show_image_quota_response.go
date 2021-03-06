package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowImageQuotaResponse struct {
	Quotas         *Quota `json:"quotas,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowImageQuotaResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowImageQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowImageQuotaResponse", string(data)}, " ")
}
