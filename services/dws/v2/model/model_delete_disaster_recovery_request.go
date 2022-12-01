package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteDisasterRecoveryRequest struct {

	// 集群的ID
	DisasterRecoveryId string `json:"disaster_recovery_id"`
}

func (o DeleteDisasterRecoveryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDisasterRecoveryRequest struct{}"
	}

	return strings.Join([]string{"DeleteDisasterRecoveryRequest", string(data)}, " ")
}
