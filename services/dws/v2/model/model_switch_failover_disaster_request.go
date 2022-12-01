package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SwitchFailoverDisasterRequest struct {

	// 容灾ID
	DisasterRecoveryId string `json:"disaster_recovery_id"`
}

func (o SwitchFailoverDisasterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchFailoverDisasterRequest struct{}"
	}

	return strings.Join([]string{"SwitchFailoverDisasterRequest", string(data)}, " ")
}
