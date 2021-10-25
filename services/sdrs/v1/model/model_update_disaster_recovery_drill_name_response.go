package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateDisasterRecoveryDrillNameResponse struct {
	DisasterRecoveryDrill *ShowDisasterRecoveryDrillParams `json:"disaster_recovery_drill,omitempty"`
	HttpStatusCode        int                              `json:"-"`
}

func (o UpdateDisasterRecoveryDrillNameResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDisasterRecoveryDrillNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateDisasterRecoveryDrillNameResponse", string(data)}, " ")
}
