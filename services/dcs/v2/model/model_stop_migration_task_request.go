package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type StopMigrationTaskRequest struct {
	// 任务ID

	TaskId string `json:"task_id"`
}

func (o StopMigrationTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StopMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"StopMigrationTaskRequest", string(data)}, " ")
}
