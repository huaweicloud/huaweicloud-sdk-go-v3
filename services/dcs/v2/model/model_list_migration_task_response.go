package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListMigrationTaskResponse struct {

	// 迁移任务数量。
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 迁移任务列表。
	MigrationTasks *[]MigrationTaskList `json:"migration_tasks,omitempty" xml:"migration_tasks"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"ListMigrationTaskResponse", string(data)}, " ")
}
