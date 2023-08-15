package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMigrationTaskResponse Response Object
type ListMigrationTaskResponse struct {

	// 迁移任务数量。
	Count *int32 `json:"count,omitempty"`

	// 迁移任务列表。
	MigrationTasks *[]MigrationTaskList `json:"migration_tasks,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"ListMigrationTaskResponse", string(data)}, " ")
}
