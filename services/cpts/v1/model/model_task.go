package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// tasks
type Task struct {

	// bench_concurrent
	BenchConcurrent *int32 `json:"bench_concurrent,omitempty" xml:"bench_concurrent"`

	// description
	Description *string `json:"description,omitempty" xml:"description"`

	// id
	Id *int32 `json:"id,omitempty" xml:"id"`

	// name
	Name *string `json:"name,omitempty" xml:"name"`

	// operate_mode
	OperateMode *int32 `json:"operate_mode,omitempty" xml:"operate_mode"`

	TaskRunInfo *TaskRunInfo `json:"task_run_info,omitempty" xml:"task_run_info"`

	// update_time
	UpdateTime *string `json:"update_time,omitempty" xml:"update_time"`
}

func (o Task) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Task struct{}"
	}

	return strings.Join([]string{"Task", string(data)}, " ")
}
