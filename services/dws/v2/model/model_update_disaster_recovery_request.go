package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDisasterRecoveryRequest struct {
	DisasterRecovery *UpdateDisasterRecoveryReq `json:"disaster_recovery,omitempty"`
}

func (o UpdateDisasterRecoveryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDisasterRecoveryRequest struct{}"
	}

	return strings.Join([]string{"UpdateDisasterRecoveryRequest", string(data)}, " ")
}
