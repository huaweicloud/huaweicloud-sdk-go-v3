package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailableZonesResponse Response Object
type ListAvailableZonesResponse struct {

	// 可用区信息
	AvailabilityZone *[]AvailabilityZones `json:"availability_zone,omitempty"`
	HttpStatusCode   int                  `json:"-"`
}

func (o ListAvailableZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableZonesResponse struct{}"
	}

	return strings.Join([]string{"ListAvailableZonesResponse", string(data)}, " ")
}
