package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量停止数据迁移任务结构体
type BatchStopMigrationTasksBody struct {
	// 数据迁移任务列表。

	MigrationTasks []string `json:"migration_tasks"`
}

func (o BatchStopMigrationTasksBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopMigrationTasksBody struct{}"
	}

	return strings.Join([]string{"BatchStopMigrationTasksBody", string(data)}, " ")
}
