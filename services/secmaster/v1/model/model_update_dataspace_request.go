package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDataspaceRequest Request Object
type UpdateDataspaceRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 数据空间ID
	DataspaceId string `json:"dataspace_id"`

	Body *UpdateDataspaceRequestBody `json:"body,omitempty"`
}

func (o UpdateDataspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataspaceRequest struct{}"
	}

	return strings.Join([]string{"UpdateDataspaceRequest", string(data)}, " ")
}
