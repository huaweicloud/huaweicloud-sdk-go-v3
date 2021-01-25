package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CinderListQuotasResponse struct {
	QuotaSet       *QuotaList `json:"quota_set,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o CinderListQuotasResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CinderListQuotasResponse struct{}"
	}

	return strings.Join([]string{"CinderListQuotasResponse", string(data)}, " ")
}
