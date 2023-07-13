package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PauseDisasterRecoveryResponse Response Object
type PauseDisasterRecoveryResponse struct {
	DisasterRecovery *DisasterRecoveryId `json:"disaster_recovery,omitempty"`
	HttpStatusCode   int                 `json:"-"`
}

func (o PauseDisasterRecoveryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseDisasterRecoveryResponse struct{}"
	}

	return strings.Join([]string{"PauseDisasterRecoveryResponse", string(data)}, " ")
}
