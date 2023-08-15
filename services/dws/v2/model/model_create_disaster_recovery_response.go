package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDisasterRecoveryResponse Response Object
type CreateDisasterRecoveryResponse struct {
	DisasterRecovery *DisasterRecoveryId `json:"disaster_recovery,omitempty"`
	HttpStatusCode   int                 `json:"-"`
}

func (o CreateDisasterRecoveryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDisasterRecoveryResponse struct{}"
	}

	return strings.Join([]string{"CreateDisasterRecoveryResponse", string(data)}, " ")
}
