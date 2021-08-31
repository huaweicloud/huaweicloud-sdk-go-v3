package model

import (
	"encoding/json"

	"strings"
)

type ShowTaskSetResqTasks struct {
	// bench_concurrent

	BenchConcurrent *int32 `json:"bench_concurrent,omitempty"`
	// description

	Description *string `json:"description,omitempty"`
	// id

	Id *int32 `json:"id,omitempty"`
	// name

	Name *string `json:"name,omitempty"`
	// operate_mode

	OperateMode *int32 `json:"operate_mode,omitempty"`

	TaskRunInfo *ShowTaskSetResqTaskRunInfo `json:"task_run_info,omitempty"`
	// update_time

	UpdateTime *string `json:"update_time,omitempty"`
}

func (o ShowTaskSetResqTasks) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTaskSetResqTasks struct{}"
	}

	return strings.Join([]string{"ShowTaskSetResqTasks", string(data)}, " ")
}
