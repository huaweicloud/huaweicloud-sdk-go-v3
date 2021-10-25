package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateDisasterRecoveryDrillRequest struct {
	Body *CreateDisasterRecoveryDrillRequestBody `json:"body,omitempty"`
}

func (o CreateDisasterRecoveryDrillRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateDisasterRecoveryDrillRequest struct{}"
	}

	return strings.Join([]string{"CreateDisasterRecoveryDrillRequest", string(data)}, " ")
}
