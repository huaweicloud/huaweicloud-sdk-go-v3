package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StopMigrationTaskRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o StopMigrationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"StopMigrationTaskRequest", string(data)}, " ")
}
