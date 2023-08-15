package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDisasterDetailRequest Request Object
type ShowDisasterDetailRequest struct {

	// 容灾ID
	DisasterRecoveryId string `json:"disaster_recovery_id"`
}

func (o ShowDisasterDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDisasterDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowDisasterDetailRequest", string(data)}, " ")
}
