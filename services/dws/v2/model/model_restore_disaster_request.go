package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreDisasterRequest Request Object
type RestoreDisasterRequest struct {

	// 容灾ID
	DisasterRecoveryId string `json:"disaster_recovery_id"`
}

func (o RestoreDisasterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreDisasterRequest struct{}"
	}

	return strings.Join([]string{"RestoreDisasterRequest", string(data)}, " ")
}
