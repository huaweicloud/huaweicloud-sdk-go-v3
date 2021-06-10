package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListCaseLimitsResponse struct {
	Config         *TenantConfigV2 `json:"config,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListCaseLimitsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCaseLimitsResponse struct{}"
	}

	return strings.Join([]string{"ListCaseLimitsResponse", string(data)}, " ")
}
