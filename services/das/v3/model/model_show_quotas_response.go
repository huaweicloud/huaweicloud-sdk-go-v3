package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowQuotasResponse struct {
	// 配额列表对象。

	Quotas         *interface{} `json:"quotas,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowQuotasResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowQuotasResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotasResponse", string(data)}, " ")
}
