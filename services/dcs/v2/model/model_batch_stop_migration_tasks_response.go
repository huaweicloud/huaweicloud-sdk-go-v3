package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchStopMigrationTasksResponse struct {

	// 数据迁移任务列表。
	MigrationTasks *[]StopMigrationTaskResult `json:"migration_tasks,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o BatchStopMigrationTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopMigrationTasksResponse struct{}"
	}

	return strings.Join([]string{"BatchStopMigrationTasksResponse", string(data)}, " ")
}
