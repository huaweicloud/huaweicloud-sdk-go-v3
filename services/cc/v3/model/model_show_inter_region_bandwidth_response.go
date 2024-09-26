package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInterRegionBandwidthResponse Response Object
type ShowInterRegionBandwidthResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	InterRegionBandwidth *InterRegionBandwidth `json:"inter_region_bandwidth"`
	HttpStatusCode       int                   `json:"-"`
}

func (o ShowInterRegionBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInterRegionBandwidthResponse struct{}"
	}

	return strings.Join([]string{"ShowInterRegionBandwidthResponse", string(data)}, " ")
}
