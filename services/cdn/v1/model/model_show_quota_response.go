package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowQuotaResponse struct {
	// 配额数组

	Quotas         *[]ShowQuotaResponseBodyQuotas `json:"quotas,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ShowQuotaResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotaResponse", string(data)}, " ")
}
