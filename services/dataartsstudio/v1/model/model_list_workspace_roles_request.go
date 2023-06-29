package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkspaceRolesRequest Request Object
type ListWorkspaceRolesRequest struct {

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 工作空间id
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o ListWorkspaceRolesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkspaceRolesRequest struct{}"
	}

	return strings.Join([]string{"ListWorkspaceRolesRequest", string(data)}, " ")
}
