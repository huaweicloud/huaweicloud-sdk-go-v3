package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowHealthMonitorRequest struct {
	HealthmonitorId string `json:"healthmonitor_id"`
}

func (o ShowHealthMonitorRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowHealthMonitorRequest struct{}"
	}

	return strings.Join([]string{"ShowHealthMonitorRequest", string(data)}, " ")
}
