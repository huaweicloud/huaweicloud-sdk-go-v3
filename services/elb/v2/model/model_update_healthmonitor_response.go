package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateHealthmonitorResponse struct {
	Healthmonitor  *HealthmonitorResp `json:"healthmonitor,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o UpdateHealthmonitorResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateHealthmonitorResponse struct{}"
	}

	return strings.Join([]string{"UpdateHealthmonitorResponse", string(data)}, " ")
}
