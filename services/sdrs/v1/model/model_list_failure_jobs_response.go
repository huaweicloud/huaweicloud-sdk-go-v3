package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListFailureJobsResponse struct {
	// 失败任务信息列表。

	FailureJobs *[]FailureJobParams `json:"failure_jobs,omitempty"`
	// 列表中失败任务个数。

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFailureJobsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListFailureJobsResponse struct{}"
	}

	return strings.Join([]string{"ListFailureJobsResponse", string(data)}, " ")
}
