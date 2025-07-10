package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLandingZoneResponse Response Object
type DeleteLandingZoneResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLandingZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLandingZoneResponse struct{}"
	}

	return strings.Join([]string{"DeleteLandingZoneResponse", string(data)}, " ")
}
