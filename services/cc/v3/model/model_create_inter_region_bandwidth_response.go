package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateInterRegionBandwidthResponse struct {
	InterRegionBandwidth *InterRegionBandwidth `json:"inter_region_bandwidth,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInterRegionBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInterRegionBandwidthResponse struct{}"
	}

	return strings.Join([]string{"CreateInterRegionBandwidthResponse", string(data)}, " ")
}
