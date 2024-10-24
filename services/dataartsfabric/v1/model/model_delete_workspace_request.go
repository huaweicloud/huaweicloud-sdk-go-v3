package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkspaceRequest Request Object
type DeleteWorkspaceRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`
}

func (o DeleteWorkspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkspaceRequest struct{}"
	}

	return strings.Join([]string{"DeleteWorkspaceRequest", string(data)}, " ")
}
