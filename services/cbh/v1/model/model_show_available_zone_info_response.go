package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailableZoneInfoResponse Response Object
type ShowAvailableZoneInfoResponse struct {
	AvailabilityZone *AvailabilityZones `json:"availability_zone,omitempty"`
	HttpStatusCode   int                `json:"-"`
}

func (o ShowAvailableZoneInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableZoneInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowAvailableZoneInfoResponse", string(data)}, " ")
}
