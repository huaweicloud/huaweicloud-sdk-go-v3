package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDisasterProgressRequest Request Object
type ShowDisasterProgressRequest struct {

	// disaster_recovery_id
	DisasterRecoveryId string `json:"disaster_recovery_id"`
}

func (o ShowDisasterProgressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDisasterProgressRequest struct{}"
	}

	return strings.Join([]string{"ShowDisasterProgressRequest", string(data)}, " ")
}
