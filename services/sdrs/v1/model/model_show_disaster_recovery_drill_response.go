package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDisasterRecoveryDrillResponse Response Object
type ShowDisasterRecoveryDrillResponse struct {
	DisasterRecoveryDrill *ShowDisasterRecoveryDrillParams `json:"disaster_recovery_drill,omitempty"`
	HttpStatusCode        int                              `json:"-"`
}

func (o ShowDisasterRecoveryDrillResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDisasterRecoveryDrillResponse struct{}"
	}

	return strings.Join([]string{"ShowDisasterRecoveryDrillResponse", string(data)}, " ")
}
