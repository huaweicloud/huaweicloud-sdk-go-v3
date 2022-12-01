package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SwitchoverDisasterRecoveryResponse struct {
	DisasterRecovery *DisasterRecoveryId `json:"disaster_recovery,omitempty"`
	HttpStatusCode   int                 `json:"-"`
}

func (o SwitchoverDisasterRecoveryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchoverDisasterRecoveryResponse struct{}"
	}

	return strings.Join([]string{"SwitchoverDisasterRecoveryResponse", string(data)}, " ")
}
