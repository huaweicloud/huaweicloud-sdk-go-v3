package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateDisasterRecoveryDrillNameRequest struct {
	// 容灾演练的ID。

	DisasterRecoveryDrillId string `json:"disaster_recovery_drill_id"`

	Body *UpdateDisasterRecoveryDrillNameRequestBody `json:"body,omitempty"`
}

func (o UpdateDisasterRecoveryDrillNameRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDisasterRecoveryDrillNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateDisasterRecoveryDrillNameRequest", string(data)}, " ")
}
