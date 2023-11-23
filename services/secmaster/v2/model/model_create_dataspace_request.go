package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataspaceRequest Request Object
type CreateDataspaceRequest struct {

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	Body *CreateDataspaceRequestBody `json:"body,omitempty"`
}

func (o CreateDataspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataspaceRequest struct{}"
	}

	return strings.Join([]string{"CreateDataspaceRequest", string(data)}, " ")
}
