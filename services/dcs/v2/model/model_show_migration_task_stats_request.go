package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowMigrationTaskStatsRequest struct {
	// 任务ID。

	TaskId string `json:"task_id"`
}

func (o ShowMigrationTaskStatsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMigrationTaskStatsRequest struct{}"
	}

	return strings.Join([]string{"ShowMigrationTaskStatsRequest", string(data)}, " ")
}
