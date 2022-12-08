package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAvailableZoneInfoResponse struct {
	AvailabilityZones *AvailabilityZones `json:"availability_zones,omitempty"`
	HttpStatusCode    int                `json:"-"`
}

func (o ShowAvailableZoneInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableZoneInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowAvailableZoneInfoResponse", string(data)}, " ")
}
