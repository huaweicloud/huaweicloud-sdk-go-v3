package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExtendedAvailabilityZonesResponse Response Object
type ListExtendedAvailabilityZonesResponse struct {
	AvailabilityZones *[]ExtendedAvailabilityZone `json:"availability_zones,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListExtendedAvailabilityZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExtendedAvailabilityZonesResponse struct{}"
	}

	return strings.Join([]string{"ListExtendedAvailabilityZonesResponse", string(data)}, " ")
}
