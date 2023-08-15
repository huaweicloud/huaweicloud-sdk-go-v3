package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResJobRequest Request Object
type UpdateResJobRequest struct {

	// 作业id
	JobId string `json:"job_id"`

	// 资源id（数据源id或场景id）
	ResourceId string `json:"resource_id"`

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	Body *UpdateResJobRequestBody `json:"body,omitempty"`
}

func (o UpdateResJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResJobRequest struct{}"
	}

	return strings.Join([]string{"UpdateResJobRequest", string(data)}, " ")
}
