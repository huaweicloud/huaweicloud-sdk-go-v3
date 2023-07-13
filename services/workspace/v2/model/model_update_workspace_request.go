package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkspaceRequest Request Object
type UpdateWorkspaceRequest struct {
	Body *ModifyWorkspaceAttributesReq `json:"body,omitempty"`
}

func (o UpdateWorkspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkspaceRequest struct{}"
	}

	return strings.Join([]string{"UpdateWorkspaceRequest", string(data)}, " ")
}
