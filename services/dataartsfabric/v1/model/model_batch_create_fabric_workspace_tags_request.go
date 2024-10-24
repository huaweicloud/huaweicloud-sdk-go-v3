package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateFabricWorkspaceTagsRequest Request Object
type BatchCreateFabricWorkspaceTagsRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *BatchCreateTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateFabricWorkspaceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateFabricWorkspaceTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateFabricWorkspaceTagsRequest", string(data)}, " ")
}
