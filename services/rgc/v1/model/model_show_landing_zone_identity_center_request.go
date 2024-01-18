package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLandingZoneIdentityCenterRequest Request Object
type ShowLandingZoneIdentityCenterRequest struct {
}

func (o ShowLandingZoneIdentityCenterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLandingZoneIdentityCenterRequest struct{}"
	}

	return strings.Join([]string{"ShowLandingZoneIdentityCenterRequest", string(data)}, " ")
}
