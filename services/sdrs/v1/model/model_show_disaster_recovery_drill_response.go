package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowDisasterRecoveryDrillResponse struct {
	DisasterRecoveryDrill *ShowDisasterRecoveryDrillParams `json:"disaster_recovery_drill,omitempty"`
	HttpStatusCode        int                              `json:"-"`
}

func (o ShowDisasterRecoveryDrillResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDisasterRecoveryDrillResponse struct{}"
	}

	return strings.Join([]string{"ShowDisasterRecoveryDrillResponse", string(data)}, " ")
}
