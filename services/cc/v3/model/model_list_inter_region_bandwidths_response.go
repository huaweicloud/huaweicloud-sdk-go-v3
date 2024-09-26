package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInterRegionBandwidthsResponse Response Object
type ListInterRegionBandwidthsResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 域间带宽实例列表。
	InterRegionBandwidths []InterRegionBandwidth `json:"inter_region_bandwidths"`
	HttpStatusCode        int                    `json:"-"`
}

func (o ListInterRegionBandwidthsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInterRegionBandwidthsResponse struct{}"
	}

	return strings.Join([]string{"ListInterRegionBandwidthsResponse", string(data)}, " ")
}
