package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowQuotaOfTenantRequest struct {
}

func (o ShowQuotaOfTenantRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowQuotaOfTenantRequest struct{}"
	}

	return strings.Join([]string{"ShowQuotaOfTenantRequest", string(data)}, " ")
}
