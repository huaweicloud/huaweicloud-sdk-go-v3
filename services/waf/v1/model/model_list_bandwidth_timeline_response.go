package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBandwidthTimelineResponse struct {

	// 带宽时间线统计数据，包括带宽（BANDWIDTH）、入带宽（IN_BANDWIDTH）以及出带宽（OUT_BANDWIDTH）统计数据。
	Body           *[]BandwidthStatisticsTimelineItem `json:"body,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListBandwidthTimelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthTimelineResponse struct{}"
	}

	return strings.Join([]string{"ListBandwidthTimelineResponse", string(data)}, " ")
}
