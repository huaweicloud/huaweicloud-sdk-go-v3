package model

import (
	"encoding/json"

	"strings"
)

//
type TaskSumbitResultBody struct {
	// 批量图像内容审核的任务标识，用于后续的结果查询。

	JobId *string `json:"job_id,omitempty"`
}

func (o TaskSumbitResultBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TaskSumbitResultBody struct{}"
	}

	return strings.Join([]string{"TaskSumbitResultBody", string(data)}, " ")
}
