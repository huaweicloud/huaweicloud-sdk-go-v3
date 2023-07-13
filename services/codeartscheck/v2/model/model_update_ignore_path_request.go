package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIgnorePathRequest Request Object
type UpdateIgnorePathRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 任务ID
	TaskId string `json:"task_id"`

	Body *UpdateIgnorePathRequestBody `json:"body,omitempty"`
}

func (o UpdateIgnorePathRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIgnorePathRequest struct{}"
	}

	return strings.Join([]string{"UpdateIgnorePathRequest", string(data)}, " ")
}
