package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInterRegionBandwidthResponse Response Object
type UpdateInterRegionBandwidthResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	InterRegionBandwidth *InterRegionBandwidth `json:"inter_region_bandwidth"`
	HttpStatusCode       int                   `json:"-"`
}

func (o UpdateInterRegionBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInterRegionBandwidthResponse struct{}"
	}

	return strings.Join([]string{"UpdateInterRegionBandwidthResponse", string(data)}, " ")
}
