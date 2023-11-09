package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailabilityZonesResponse Response Object
type ListAvailabilityZonesResponse struct {
	AvailabilityZones *AvailabilityZones `json:"availability_zones,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAvailabilityZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailabilityZonesResponse struct{}"
	}

	return strings.Join([]string{"ListAvailabilityZonesResponse", string(data)}, " ")
}
