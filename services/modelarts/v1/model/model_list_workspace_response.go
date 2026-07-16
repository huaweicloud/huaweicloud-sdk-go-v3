package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkspaceResponse Response Object
type ListWorkspaceResponse struct {

	// 工作空间的总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 此次请求返回的工作空间个数。
	Count *int32 `json:"count,omitempty"`

	// workspace属性列表。
	Workspaces     *[]WorkspaceResponse `json:"workspaces,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListWorkspaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkspaceResponse struct{}"
	}

	return strings.Join([]string{"ListWorkspaceResponse", string(data)}, " ")
}
