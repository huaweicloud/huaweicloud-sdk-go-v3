package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListQpsTimelineResponse struct {

	// 安全统计的时间线
	Body           *[]StatisticsTimelineItem `json:"body,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListQpsTimelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQpsTimelineResponse struct{}"
	}

	return strings.Join([]string{"ListQpsTimelineResponse", string(data)}, " ")
}
