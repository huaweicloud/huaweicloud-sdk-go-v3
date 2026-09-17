package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTaskRequest Request Object
type CreateTaskRequest struct {

	// 项目id
	ProjectUuid string `json:"project_uuid"`

	Body *TaskInfo `json:"body,omitempty"`
}

func (o CreateTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateTaskRequest", string(data)}, " ")
}
