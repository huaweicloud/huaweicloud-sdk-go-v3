package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDomainQuotaRequest struct {
	// 账号ID。

	DomainId string `json:"domain_id"`
}

func (o ShowDomainQuotaRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDomainQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainQuotaRequest", string(data)}, " ")
}
