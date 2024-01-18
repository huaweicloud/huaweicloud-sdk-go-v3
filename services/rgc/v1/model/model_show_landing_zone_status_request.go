package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLandingZoneStatusRequest Request Object
type ShowLandingZoneStatusRequest struct {
}

func (o ShowLandingZoneStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLandingZoneStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowLandingZoneStatusRequest", string(data)}, " ")
}
