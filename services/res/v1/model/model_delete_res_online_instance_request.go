package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteResOnlineInstanceRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`

	// 资源id（数据源id或场景id）。
	ResourceId string `json:"resource_id"`

	// 作业id。
	JobId string `json:"job_id"`
}

func (o DeleteResOnlineInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResOnlineInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteResOnlineInstanceRequest", string(data)}, " ")
}
