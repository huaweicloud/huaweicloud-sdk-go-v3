package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowVodStatisticsResponse struct {
	// 统计起始时间

	StartTime *string `json:"start_time,omitempty"`
	// 采样时间间隔

	Interval *int32 `json:"interval,omitempty"`

	SampleData     *[]VodSampleData `json:"sample_data,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowVodStatisticsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowVodStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowVodStatisticsResponse", string(data)}, " ")
}
