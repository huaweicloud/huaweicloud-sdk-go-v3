package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowProjectQuotaRequest struct {
	ProjectId string `json:"project_id"`
}

func (o ShowProjectQuotaRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProjectQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectQuotaRequest", string(data)}, " ")
}
