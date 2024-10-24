package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkspaceRequest Request Object
type CreateWorkspaceRequest struct {
	Body *CreateWorkspaceRequestInput `json:"body,omitempty"`
}

func (o CreateWorkspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkspaceRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkspaceRequest", string(data)}, " ")
}
