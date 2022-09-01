package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SetOnlineMigrationTaskRequest struct {

	// 在线迁移任务ID。
	TaskId string `json:"task_id" xml:"task_id"`

	Body *SetOnlineMigrationTaskBody `json:"body,omitempty" xml:"body"`
}

func (o SetOnlineMigrationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetOnlineMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"SetOnlineMigrationTaskRequest", string(data)}, " ")
}
