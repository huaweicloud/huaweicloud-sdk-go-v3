package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLandingZoneRequest Request Object
type DeleteLandingZoneRequest struct {
}

func (o DeleteLandingZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLandingZoneRequest struct{}"
	}

	return strings.Join([]string{"DeleteLandingZoneRequest", string(data)}, " ")
}
