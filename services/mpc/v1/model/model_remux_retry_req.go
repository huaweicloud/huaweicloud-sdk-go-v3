package model

import (
	"encoding/json"

	"strings"
)

type RemuxRetryReq struct {
	// 任务Id。

	TaskId *string `json:"task_id,omitempty"`
}

func (o RemuxRetryReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RemuxRetryReq struct{}"
	}

	return strings.Join([]string{"RemuxRetryReq", string(data)}, " ")
}
