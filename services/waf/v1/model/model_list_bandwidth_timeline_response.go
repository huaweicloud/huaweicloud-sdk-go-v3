package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBandwidthTimelineResponse struct {

	// 安全统计的带宽时间线数据
	Body           *[]BandwidthStatisticsTimelineItem `json:"body,omitempty" xml:"body"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListBandwidthTimelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthTimelineResponse struct{}"
	}

	return strings.Join([]string{"ListBandwidthTimelineResponse", string(data)}, " ")
}
