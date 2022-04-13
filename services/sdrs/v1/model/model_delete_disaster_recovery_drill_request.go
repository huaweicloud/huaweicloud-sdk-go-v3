package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteDisasterRecoveryDrillRequest struct {
	// 容灾演练的ID。

	DisasterRecoveryDrillId string `json:"disaster_recovery_drill_id"`
}

func (o DeleteDisasterRecoveryDrillRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDisasterRecoveryDrillRequest struct{}"
	}

	return strings.Join([]string{"DeleteDisasterRecoveryDrillRequest", string(data)}, " ")
}
