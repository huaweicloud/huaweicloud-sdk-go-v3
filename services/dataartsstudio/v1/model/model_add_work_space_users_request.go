package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddWorkSpaceUsersRequest Request Object
type AddWorkSpaceUsersRequest struct {

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	Body *ApigWorkspaceUserDto `json:"body,omitempty"`
}

func (o AddWorkSpaceUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddWorkSpaceUsersRequest struct{}"
	}

	return strings.Join([]string{"AddWorkSpaceUsersRequest", string(data)}, " ")
}
