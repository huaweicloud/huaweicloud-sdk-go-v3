package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShareTypesAttribution struct {
	Capacity *ShareTypesAttributionCapacity `json:"capacity,omitempty"`

	Bandwidth *ShareTypesAttributionBandwidth `json:"bandwidth,omitempty"`

	Iops *ShareTypesAttributionIops `json:"iops,omitempty"`

	SingleChannel4kLatency *ShareTypesAttributionSingleChannel4KLatency `json:"single_channel_4k_latency,omitempty"`
}

func (o ShareTypesAttribution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareTypesAttribution struct{}"
	}

	return strings.Join([]string{"ShareTypesAttribution", string(data)}, " ")
}
