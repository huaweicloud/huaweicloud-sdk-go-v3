package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkspaceusersRequest Request Object
type DeleteWorkspaceusersRequest struct {

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	Body *ApigDelUserParams `json:"body,omitempty"`
}

func (o DeleteWorkspaceusersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkspaceusersRequest struct{}"
	}

	return strings.Join([]string{"DeleteWorkspaceusersRequest", string(data)}, " ")
}
