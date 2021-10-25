package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDisasterRecoveryDrillRequest struct {
	// 容灾演练的ID。

	DisasterRecoveryDrillId string `json:"disaster_recovery_drill_id"`
}

func (o ShowDisasterRecoveryDrillRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDisasterRecoveryDrillRequest struct{}"
	}

	return strings.Join([]string{"ShowDisasterRecoveryDrillRequest", string(data)}, " ")
}
