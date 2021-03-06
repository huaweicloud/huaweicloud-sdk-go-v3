package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowHealthmonitorsResponse struct {
	Healthmonitor  *HealthmonitorResp `json:"healthmonitor,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowHealthmonitorsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowHealthmonitorsResponse struct{}"
	}

	return strings.Join([]string{"ShowHealthmonitorsResponse", string(data)}, " ")
}
