package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowEnterpriseProjectQuotaRequest struct {
}

func (o ShowEnterpriseProjectQuotaRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowEnterpriseProjectQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowEnterpriseProjectQuotaRequest", string(data)}, " ")
}
