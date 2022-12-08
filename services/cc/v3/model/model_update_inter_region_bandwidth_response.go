package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateInterRegionBandwidthResponse struct {
	InterRegionBandwidth *InterRegionBandwidth `json:"inter_region_bandwidth,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateInterRegionBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInterRegionBandwidthResponse struct{}"
	}

	return strings.Join([]string{"UpdateInterRegionBandwidthResponse", string(data)}, " ")
}
