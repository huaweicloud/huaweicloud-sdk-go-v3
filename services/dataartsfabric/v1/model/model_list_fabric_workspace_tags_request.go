package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFabricWorkspaceTagsRequest Request Object
type ListFabricWorkspaceTagsRequest struct {

	// Workspace的ID
	WorkspaceId string `json:"workspace_id"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数默认为1000，limit最多为1000,不能为负数，最小值为1
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListFabricWorkspaceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFabricWorkspaceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListFabricWorkspaceTagsRequest", string(data)}, " ")
}
