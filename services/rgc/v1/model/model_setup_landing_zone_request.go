package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetupLandingZoneRequest Request Object
type SetupLandingZoneRequest struct {
	Body *SetupLandingZoneReqBody `json:"body,omitempty"`
}

func (o SetupLandingZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetupLandingZoneRequest struct{}"
	}

	return strings.Join([]string{"SetupLandingZoneRequest", string(data)}, " ")
}
