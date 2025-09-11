package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryFlavorCapacityAzInfo struct {
	RegionId *string `json:"region_id,omitempty"`

	AvailabilityZone *string `json:"availability_zone,omitempty"`

	Prefer *bool `json:"prefer,omitempty"`
}

func (o QueryFlavorCapacityAzInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryFlavorCapacityAzInfo struct{}"
	}

	return strings.Join([]string{"QueryFlavorCapacityAzInfo", string(data)}, " ")
}
