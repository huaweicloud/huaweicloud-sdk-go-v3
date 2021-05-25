package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowMigrationTaskRequest struct {
	// 任务ID。

	TaskId string `json:"task_id"`
}

func (o ShowMigrationTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowMigrationTaskRequest", string(data)}, " ")
}
