package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResWorkspaceRequest Request Object
type DeleteResWorkspaceRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`
}

func (o DeleteResWorkspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResWorkspaceRequest struct{}"
	}

	return strings.Join([]string{"DeleteResWorkspaceRequest", string(data)}, " ")
}
