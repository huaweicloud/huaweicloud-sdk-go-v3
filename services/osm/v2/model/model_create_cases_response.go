package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateCasesResponse struct {
	// 工单id

	IncidentId     *string `json:"incident_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCasesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCasesResponse struct{}"
	}

	return strings.Join([]string{"CreateCasesResponse", string(data)}, " ")
}
