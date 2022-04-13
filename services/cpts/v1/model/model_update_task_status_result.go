package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTaskStatusResult struct {
	// task_run_id

	TaskRunId *int32 `json:"task_run_id,omitempty"`
}

func (o UpdateTaskStatusResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskStatusResult struct{}"
	}

	return strings.Join([]string{"UpdateTaskStatusResult", string(data)}, " ")
}
