package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchStartJobsResponse struct {
	// 批量启动实时灾备任务返回列表。

	Results *[]StartJobResp `json:"results,omitempty"`
	// 总数。

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchStartJobsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchStartJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchStartJobsResponse", string(data)}, " ")
}
