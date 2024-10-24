package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkspacesResponse Response Object
type ListWorkspacesResponse struct {

	// workspace总数
	Total *int32 `json:"total,omitempty"`

	// workaspace列表
	Workspaces *[]Workspace `json:"workspaces,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListWorkspacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkspacesResponse struct{}"
	}

	return strings.Join([]string{"ListWorkspacesResponse", string(data)}, " ")
}
