package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDisasterRecoveryDrillRequest struct {
	Body *CreateDisasterRecoveryDrillRequestBody `json:"body,omitempty"`
}

func (o CreateDisasterRecoveryDrillRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDisasterRecoveryDrillRequest struct{}"
	}

	return strings.Join([]string{"CreateDisasterRecoveryDrillRequest", string(data)}, " ")
}
