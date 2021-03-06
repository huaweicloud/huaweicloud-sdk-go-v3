package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowQuotaDefaultsRequest struct {
}

func (o ShowQuotaDefaultsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowQuotaDefaultsRequest struct{}"
	}

	return strings.Join([]string{"ShowQuotaDefaultsRequest", string(data)}, " ")
}
