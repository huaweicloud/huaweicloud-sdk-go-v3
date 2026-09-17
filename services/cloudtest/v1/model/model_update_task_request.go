package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaskRequest Request Object
type UpdateTaskRequest struct {

	// 项目id
	ProjectUuid string `json:"project_uuid"`

	// 测试套件uri
	TaskUri string `json:"task_uri"`

	Body *TaskInfo `json:"body,omitempty"`
}

func (o UpdateTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskRequest struct{}"
	}

	return strings.Join([]string{"UpdateTaskRequest", string(data)}, " ")
}
