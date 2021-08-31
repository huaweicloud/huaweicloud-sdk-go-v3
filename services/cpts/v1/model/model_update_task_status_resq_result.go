package model

import (
	"encoding/json"

	"strings"
)

// result
type UpdateTaskStatusResqResult struct {
	// task_run_id

	TaskRunId *int32 `json:"task_run_id,omitempty"`
}

func (o UpdateTaskStatusResqResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTaskStatusResqResult struct{}"
	}

	return strings.Join([]string{"UpdateTaskStatusResqResult", string(data)}, " ")
}
