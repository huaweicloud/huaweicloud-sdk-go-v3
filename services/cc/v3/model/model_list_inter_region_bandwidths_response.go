package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListInterRegionBandwidthsResponse struct {

	// 域间带宽实例列表。
	InterRegionBandwidth *[]InterRegionBandwidth `json:"inter_region_bandwidth,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListInterRegionBandwidthsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInterRegionBandwidthsResponse struct{}"
	}

	return strings.Join([]string{"ListInterRegionBandwidthsResponse", string(data)}, " ")
}
