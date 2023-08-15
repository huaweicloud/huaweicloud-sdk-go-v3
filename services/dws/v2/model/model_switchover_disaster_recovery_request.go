package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchoverDisasterRecoveryRequest Request Object
type SwitchoverDisasterRecoveryRequest struct {

	// 容灾ID
	DisasterRecoveryId string `json:"disaster_recovery_id"`
}

func (o SwitchoverDisasterRecoveryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchoverDisasterRecoveryRequest struct{}"
	}

	return strings.Join([]string{"SwitchoverDisasterRecoveryRequest", string(data)}, " ")
}
