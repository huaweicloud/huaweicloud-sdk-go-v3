package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResWorkspaceRequest Request Object
type UpdateResWorkspaceRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`

	Body *UpdateResWorkspaceRequestBody `json:"body,omitempty"`
}

func (o UpdateResWorkspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResWorkspaceRequest struct{}"
	}

	return strings.Join([]string{"UpdateResWorkspaceRequest", string(data)}, " ")
}
