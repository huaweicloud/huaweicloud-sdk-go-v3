package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInterRegionBandwidthResponse Response Object
type ShowInterRegionBandwidthResponse struct {
	InterRegionBandwidth *InterRegionBandwidth `json:"inter_region_bandwidth,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInterRegionBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInterRegionBandwidthResponse struct{}"
	}

	return strings.Join([]string{"ShowInterRegionBandwidthResponse", string(data)}, " ")
}
