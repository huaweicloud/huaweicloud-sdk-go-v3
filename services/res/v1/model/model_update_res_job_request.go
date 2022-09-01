package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateResJobRequest struct {

	// 作业id
	JobId string `json:"job_id" xml:"job_id"`

	// 资源id（数据源id或场景id）
	ResourceId string `json:"resource_id" xml:"resource_id"`

	// 工作空间id
	WorkspaceId string `json:"workspace_id" xml:"workspace_id"`

	Body *UpdateResJobRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateResJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResJobRequest struct{}"
	}

	return strings.Join([]string{"UpdateResJobRequest", string(data)}, " ")
}
