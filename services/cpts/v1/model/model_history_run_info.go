package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HistoryRunInfo struct {

	// name
	Name *string `json:"name,omitempty"`

	// run_id
	RunId *float64 `json:"run_id,omitempty"`

	// run_type
	RunType *float64 `json:"run_type,omitempty"`

	// start_time
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// continue_time
	ContinueTime *float64 `json:"continue_time,omitempty"`

	// temp_names
	TempNames *[]TempName `json:"temp_names,omitempty"`

	// 任务间用例是否并行执行
	Parallel *bool `json:"parallel,omitempty"`
}

func (o HistoryRunInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryRunInfo struct{}"
	}

	return strings.Join([]string{"HistoryRunInfo", string(data)}, " ")
}
