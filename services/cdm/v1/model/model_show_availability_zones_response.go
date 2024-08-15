package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailabilityZonesResponse Response Object
type ShowAvailabilityZonesResponse struct {

	// 区域ID。
	RegionId *string `json:"regionId,omitempty"`

	// 默认可用区。
	DefaultAZ *string `json:"defaultAZ,omitempty"`

	// 可用区。
	AvailableZones *[]CdmClusterAvailabilityZone `json:"availableZones,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowAvailabilityZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailabilityZonesResponse struct{}"
	}

	return strings.Join([]string{"ShowAvailabilityZonesResponse", string(data)}, " ")
}
