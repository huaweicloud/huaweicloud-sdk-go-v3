package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkspaceResponse Response Object
type UpdateWorkspaceResponse struct {

	// 工作空间ID。
	WorkspaceId    *string `json:"workspace_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateWorkspaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkspaceResponse struct{}"
	}

	return strings.Join([]string{"UpdateWorkspaceResponse", string(data)}, " ")
}
