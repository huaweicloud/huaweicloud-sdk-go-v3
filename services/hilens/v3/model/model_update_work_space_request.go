package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkSpaceRequest Request Object
type UpdateWorkSpaceRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *UpdateDescription `json:"body,omitempty"`
}

func (o UpdateWorkSpaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkSpaceRequest struct{}"
	}

	return strings.Join([]string{"UpdateWorkSpaceRequest", string(data)}, " ")
}
