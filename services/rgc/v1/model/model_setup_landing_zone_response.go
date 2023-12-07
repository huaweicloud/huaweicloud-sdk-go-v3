package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetupLandingZoneResponse Response Object
type SetupLandingZoneResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetupLandingZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetupLandingZoneResponse struct{}"
	}

	return strings.Join([]string{"SetupLandingZoneResponse", string(data)}, " ")
}
