package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartDisasterRecoveryRequest Request Object
type StartDisasterRecoveryRequest struct {

	// 容灾ID
	DisasterRecoveryId string `json:"disaster_recovery_id"`
}

func (o StartDisasterRecoveryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartDisasterRecoveryRequest struct{}"
	}

	return strings.Join([]string{"StartDisasterRecoveryRequest", string(data)}, " ")
}
