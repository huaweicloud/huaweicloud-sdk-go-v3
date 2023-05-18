package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateDisasterInfoResponse struct {
	DisasterRecovery *DisasterRecoveryId `json:"disaster_recovery,omitempty"`
	HttpStatusCode   int                 `json:"-"`
}

func (o UpdateDisasterInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDisasterInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateDisasterInfoResponse", string(data)}, " ")
}
