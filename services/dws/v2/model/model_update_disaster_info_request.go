package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDisasterInfoRequest Request Object
type UpdateDisasterInfoRequest struct {

	// 容灾ID
	DisasterRecoveryId string `json:"disaster_recovery_id"`

	Body *UpdateDisasterRecoveryRequest `json:"body,omitempty"`
}

func (o UpdateDisasterInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDisasterInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateDisasterInfoRequest", string(data)}, " ")
}
