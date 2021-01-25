package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdateHealthMonitorRequestBody struct {
	Healthmonitor *UpdateHealthMonitorOption `json:"healthmonitor"`
}

func (o UpdateHealthMonitorRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateHealthMonitorRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHealthMonitorRequestBody", string(data)}, " ")
}
