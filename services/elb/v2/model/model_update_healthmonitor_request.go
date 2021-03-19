package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateHealthmonitorRequest struct {
	HealthmonitorId string `json:"healthmonitor_id"`

	Body *UpdateHealthmonitorRequestBody `json:"body,omitempty"`
}

func (o UpdateHealthmonitorRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateHealthmonitorRequest struct{}"
	}

	return strings.Join([]string{"UpdateHealthmonitorRequest", string(data)}, " ")
}
