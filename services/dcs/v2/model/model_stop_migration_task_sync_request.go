package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StopMigrationTaskSyncRequest struct {

	// 任务ID
	TaskId string `json:"task_id" xml:"task_id"`
}

func (o StopMigrationTaskSyncRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopMigrationTaskSyncRequest struct{}"
	}

	return strings.Join([]string{"StopMigrationTaskSyncRequest", string(data)}, " ")
}
