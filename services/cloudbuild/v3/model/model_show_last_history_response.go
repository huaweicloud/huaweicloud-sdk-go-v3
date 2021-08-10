package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowLastHistoryResponse struct {
	// 构建记录id--唯一key

	RecordId *string `json:"record_id,omitempty"`
	// 构建任务ID

	JobId *string `json:"job_id,omitempty"`
	// 构建任务名称

	JobName *string `json:"job_name,omitempty"`
	// 构建编号

	BuildNumber *int32 `json:"build_number,omitempty"`
	// 构建开始时间

	StartTime *string `json:"start_time,omitempty"`
	// 构建结束时间

	EndTime *string `json:"end_time,omitempty"`
	// 构建执行结果

	Result *string `json:"result,omitempty"`
	// commitId

	CommitId       *string `json:"commit_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLastHistoryResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowLastHistoryResponse struct{}"
	}

	return strings.Join([]string{"ShowLastHistoryResponse", string(data)}, " ")
}
