package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListP2cVgwAvailabilityZonesResponse Response Object
type ListP2cVgwAvailabilityZonesResponse struct {

	// 可用区列表
	AvailabilityZones *[]string `json:"availability_zones,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListP2cVgwAvailabilityZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListP2cVgwAvailabilityZonesResponse struct{}"
	}

	return strings.Join([]string{"ListP2cVgwAvailabilityZonesResponse", string(data)}, " ")
}
