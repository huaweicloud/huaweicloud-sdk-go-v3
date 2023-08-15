package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailabilityZoneResponse Response Object
type ListAvailabilityZoneResponse struct {

	// 可用区列表
	AvailabilityZones *[]AvailableZone `json:"availability_zones,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAvailabilityZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailabilityZoneResponse struct{}"
	}

	return strings.Join([]string{"ListAvailabilityZoneResponse", string(data)}, " ")
}
