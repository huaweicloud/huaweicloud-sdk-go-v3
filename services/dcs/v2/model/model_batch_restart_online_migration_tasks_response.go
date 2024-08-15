package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRestartOnlineMigrationTasksResponse Response Object
type BatchRestartOnlineMigrationTasksResponse struct {

	// 数据迁移任务列表。
	MigrationTasks *[]BatchRestartMigrationTaskResult `json:"migration_tasks,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o BatchRestartOnlineMigrationTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRestartOnlineMigrationTasksResponse struct{}"
	}

	return strings.Join([]string{"BatchRestartOnlineMigrationTasksResponse", string(data)}, " ")
}
