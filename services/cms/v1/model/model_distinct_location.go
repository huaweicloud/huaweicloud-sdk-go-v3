package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type DistinctLocation struct {

	// region的id
	RegionId string `json:"region_id"`

	// az的id
	AvailabilityZoneId *string `json:"availability_zone_id,omitempty"`
}

func (o DistinctLocation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DistinctLocation struct{}"
	}

	return strings.Join([]string{"DistinctLocation", string(data)}, " ")
}
