package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResOnlineInstanceRequest Request Object
type UpdateResOnlineInstanceRequest struct {

	// 作业id。
	JobId string `json:"job_id"`

	// 资源id（数据源id或场景id）。
	ResourceId string `json:"resource_id"`

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`

	Body *UpdateResOnlineInstanceRequestBody `json:"body,omitempty"`
}

func (o UpdateResOnlineInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResOnlineInstanceRequest struct{}"
	}

	return strings.Join([]string{"UpdateResOnlineInstanceRequest", string(data)}, " ")
}
