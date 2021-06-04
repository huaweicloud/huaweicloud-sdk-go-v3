package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowQuotasRequest struct {
}

func (o ShowQuotasRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowQuotasRequest", string(data)}, " ")
}
