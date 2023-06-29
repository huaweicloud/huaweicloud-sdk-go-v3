package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkSpaceUserOrGroupRequest Request Object
type UpdateWorkSpaceUserOrGroupRequest struct {

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	// 用户id
	UserId string `json:"user_id"`

	Body *ApigWorkspaceUserDto `json:"body,omitempty"`
}

func (o UpdateWorkSpaceUserOrGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkSpaceUserOrGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateWorkSpaceUserOrGroupRequest", string(data)}, " ")
}
