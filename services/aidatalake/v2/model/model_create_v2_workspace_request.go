package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateV2WorkspaceRequest Request Object
type CreateV2WorkspaceRequest struct {
	Body *CreateV2WorkspaceRequestInput `json:"body,omitempty"`
}

func (o CreateV2WorkspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateV2WorkspaceRequest struct{}"
	}

	return strings.Join([]string{"CreateV2WorkspaceRequest", string(data)}, " ")
}
