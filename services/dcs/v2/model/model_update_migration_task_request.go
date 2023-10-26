package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMigrationTaskRequest Request Object
type UpdateMigrationTaskRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`

	Body *MigrationUpdateRequestEntity `json:"body,omitempty"`
}

func (o UpdateMigrationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"UpdateMigrationTaskRequest", string(data)}, " ")
}
