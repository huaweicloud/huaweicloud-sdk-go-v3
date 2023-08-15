package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResOnlineInstanceRequest Request Object
type CreateResOnlineInstanceRequest struct {

	// 资源id（数据源id或场景id）。
	ResourceId string `json:"resource_id"`

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`

	Body *CreateResOnlineInstanceRequestBody `json:"body,omitempty"`
}

func (o CreateResOnlineInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResOnlineInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateResOnlineInstanceRequest", string(data)}, " ")
}
