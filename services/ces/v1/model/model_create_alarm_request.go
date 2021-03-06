package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateAlarmRequest struct {
	Body *CreateAlarmRequestBody `json:"body,omitempty"`
}

func (o CreateAlarmRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAlarmRequest struct{}"
	}

	return strings.Join([]string{"CreateAlarmRequest", string(data)}, " ")
}
