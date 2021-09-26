package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListResponseCodeTimelineResponse struct {
	// 安全统计的时间线

	Body           *[]StatisticsTimelineItem `json:"body,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListResponseCodeTimelineResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListResponseCodeTimelineResponse struct{}"
	}

	return strings.Join([]string{"ListResponseCodeTimelineResponse", string(data)}, " ")
}
