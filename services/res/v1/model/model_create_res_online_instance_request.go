package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateResOnlineInstanceRequest struct {

	// 资源id（数据源id或场景id）。
	ResourceId string `json:"resource_id" xml:"resource_id"`

	// 工作空间id。
	WorkspaceId string `json:"workspace_id" xml:"workspace_id"`

	Body *CreateResOnlineInstanceRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateResOnlineInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResOnlineInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateResOnlineInstanceRequest", string(data)}, " ")
}
