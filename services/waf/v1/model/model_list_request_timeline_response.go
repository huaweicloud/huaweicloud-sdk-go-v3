package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRequestTimelineResponse Response Object
type ListRequestTimelineResponse struct {

	// 安全统计的请求时间线数据
	Body           *[]StatisticsTimelineItem `json:"body,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListRequestTimelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRequestTimelineResponse struct{}"
	}

	return strings.Join([]string{"ListRequestTimelineResponse", string(data)}, " ")
}
