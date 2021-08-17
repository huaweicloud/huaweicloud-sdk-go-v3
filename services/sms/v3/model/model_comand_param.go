package model

import (
	"encoding/json"

	"strings"
)

// 命令响应参数
type ComandParam struct {
	// 任务id

	TaskId *string `json:"task_id,omitempty"`
	// 桶名

	Bucket *string `json:"bucket,omitempty"`
}

func (o ComandParam) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ComandParam struct{}"
	}

	return strings.Join([]string{"ComandParam", string(data)}, " ")
}
