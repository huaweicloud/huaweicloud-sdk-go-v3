package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRestartOnlineMigrationTasksBody 批量重启在线迁移任务结构体
type BatchRestartOnlineMigrationTasksBody struct {

	// 数据迁移任务列表。
	MigrationTasks []string `json:"migration_tasks"`
}

func (o BatchRestartOnlineMigrationTasksBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRestartOnlineMigrationTasksBody struct{}"
	}

	return strings.Join([]string{"BatchRestartOnlineMigrationTasksBody", string(data)}, " ")
}
