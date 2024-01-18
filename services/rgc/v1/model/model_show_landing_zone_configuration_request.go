package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLandingZoneConfigurationRequest Request Object
type ShowLandingZoneConfigurationRequest struct {
}

func (o ShowLandingZoneConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLandingZoneConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ShowLandingZoneConfigurationRequest", string(data)}, " ")
}
