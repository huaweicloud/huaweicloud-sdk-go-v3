package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateTaskSettingsRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 任务ID
	TaskId string `json:"task_id"`

	Body *UpdateTaskSettingsRequestBody `json:"body,omitempty"`
}

func (o UpdateTaskSettingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskSettingsRequest struct{}"
	}

	return strings.Join([]string{"UpdateTaskSettingsRequest", string(data)}, " ")
}
