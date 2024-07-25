package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkspacesResponse Response Object
type ListWorkspacesResponse struct {

	// 空间信息
	Workspaces *[]CreateWorkspaceResponseBody `json:"workspaces,omitempty"`

	// 数据总量
	Count          float32 `json:"count,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListWorkspacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkspacesResponse struct{}"
	}

	return strings.Join([]string{"ListWorkspacesResponse", string(data)}, " ")
}
