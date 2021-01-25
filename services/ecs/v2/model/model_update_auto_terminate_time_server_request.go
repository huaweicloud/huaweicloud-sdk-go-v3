package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateAutoTerminateTimeServerRequest struct {
	ServerId string                                    `json:"server_id"`
	Body     *UpdateAutoTerminateTimeServerRequestBody `json:"body,omitempty"`
}

func (o UpdateAutoTerminateTimeServerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAutoTerminateTimeServerRequest struct{}"
	}

	return strings.Join([]string{"UpdateAutoTerminateTimeServerRequest", string(data)}, " ")
}
