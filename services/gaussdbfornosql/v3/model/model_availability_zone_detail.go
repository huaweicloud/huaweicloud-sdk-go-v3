package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AvailabilityZoneDetail 主备版的多AZ可用区详情
type AvailabilityZoneDetail struct {

	// 主可用区，应为单可用区且和备可用区不同
	PrimaryAvailabilityZone string `json:"primary_availability_zone"`

	// 备可用区，应为单可用区且和主可用区不同
	SecondaryAvailabilityZone string `json:"secondary_availability_zone"`
}

func (o AvailabilityZoneDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailabilityZoneDetail struct{}"
	}

	return strings.Join([]string{"AvailabilityZoneDetail", string(data)}, " ")
}
