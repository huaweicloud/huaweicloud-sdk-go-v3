package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkspaceRequest Request Object
type UpdateWorkspaceRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *WorkspaceVo `json:"body,omitempty"`
}

func (o UpdateWorkspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkspaceRequest struct{}"
	}

	return strings.Join([]string{"UpdateWorkspaceRequest", string(data)}, " ")
}
