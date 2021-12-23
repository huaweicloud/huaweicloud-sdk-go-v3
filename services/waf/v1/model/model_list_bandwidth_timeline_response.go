package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBandwidthTimelineResponse struct {
	// 安全统计的时间线

	Body           *[]StatisticsTimelineItem `json:"body,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListBandwidthTimelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthTimelineResponse struct{}"
	}

	return strings.Join([]string{"ListBandwidthTimelineResponse", string(data)}, " ")
}
