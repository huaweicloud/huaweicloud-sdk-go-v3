package model

import (
	"encoding/json"

	"strings"
)

type TaskStatus struct {
	// 任务ID，供调用者查询任务进度。

	JobID *string `json:"jobID,omitempty"`
}

func (o TaskStatus) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TaskStatus struct{}"
	}

	return strings.Join([]string{"TaskStatus", string(data)}, " ")
}
