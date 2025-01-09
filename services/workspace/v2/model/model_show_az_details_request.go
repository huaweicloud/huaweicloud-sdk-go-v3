package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAzDetailsRequest Request Object
type ShowAzDetailsRequest struct {

	// 可用分区ID。
	AvailabilityZoneId string `json:"availability_zone_id"`
}

func (o ShowAzDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAzDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowAzDetailsRequest", string(data)}, " ")
}
