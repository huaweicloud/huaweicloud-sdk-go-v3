package model

import (
	"encoding/json"

	"strings"
)

type CommonCreateTaskRsp struct {
	// 任务ID

	TaskId *string `json:"task_id,omitempty"`
}

func (o CommonCreateTaskRsp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CommonCreateTaskRsp struct{}"
	}

	return strings.Join([]string{"CommonCreateTaskRsp", string(data)}, " ")
}
