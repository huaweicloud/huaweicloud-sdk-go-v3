package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateHealthMonitorRequestBody struct {
	Healthmonitor *CreateHealthMonitorOption `json:"healthmonitor"`
}

func (o CreateHealthMonitorRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateHealthMonitorRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHealthMonitorRequestBody", string(data)}, " ")
}
