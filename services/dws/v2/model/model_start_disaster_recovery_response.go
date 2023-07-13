package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartDisasterRecoveryResponse Response Object
type StartDisasterRecoveryResponse struct {
	DisasterRecovery *DisasterRecoveryId `json:"disaster_recovery,omitempty"`
	HttpStatusCode   int                 `json:"-"`
}

func (o StartDisasterRecoveryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartDisasterRecoveryResponse struct{}"
	}

	return strings.Join([]string{"StartDisasterRecoveryResponse", string(data)}, " ")
}
