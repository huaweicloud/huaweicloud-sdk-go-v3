package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowCdnStatisticsResponse struct {
	// 统计起始时间

	StartTime *string `json:"start_time,omitempty"`
	// 采样时间间隔

	Interval *int32 `json:"interval,omitempty"`

	Values         *[]int64 `json:"values,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowCdnStatisticsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowCdnStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowCdnStatisticsResponse", string(data)}, " ")
}
