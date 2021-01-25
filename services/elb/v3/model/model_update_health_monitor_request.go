package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateHealthMonitorRequest struct {
	HealthmonitorId string                          `json:"healthmonitor_id"`
	Body            *UpdateHealthMonitorRequestBody `json:"body,omitempty"`
}

func (o UpdateHealthMonitorRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateHealthMonitorRequest struct{}"
	}

	return strings.Join([]string{"UpdateHealthMonitorRequest", string(data)}, " ")
}
