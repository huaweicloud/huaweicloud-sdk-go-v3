package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PauseDisasterRecoveryRequest Request Object
type PauseDisasterRecoveryRequest struct {

	// 容灾ID
	DisasterRecoveryId string `json:"disaster_recovery_id"`
}

func (o PauseDisasterRecoveryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseDisasterRecoveryRequest struct{}"
	}

	return strings.Join([]string{"PauseDisasterRecoveryRequest", string(data)}, " ")
}
