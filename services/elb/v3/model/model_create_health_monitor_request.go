package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateHealthMonitorRequest struct {
	Body *CreateHealthMonitorRequestBody `json:"body,omitempty"`
}

func (o CreateHealthMonitorRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateHealthMonitorRequest struct{}"
	}

	return strings.Join([]string{"CreateHealthMonitorRequest", string(data)}, " ")
}
