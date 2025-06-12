package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkspaceRequest Request Object
type ShowWorkspaceRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`
}

func (o ShowWorkspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkspaceRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkspaceRequest", string(data)}, " ")
}
