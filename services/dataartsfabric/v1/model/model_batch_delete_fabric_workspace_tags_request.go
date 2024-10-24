package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteFabricWorkspaceTagsRequest Request Object
type BatchDeleteFabricWorkspaceTagsRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *BatchDeleteTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteFabricWorkspaceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteFabricWorkspaceTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteFabricWorkspaceTagsRequest", string(data)}, " ")
}
