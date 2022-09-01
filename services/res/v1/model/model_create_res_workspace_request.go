package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateResWorkspaceRequest struct {
	Body *CreateResWorkspaceRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateResWorkspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResWorkspaceRequest struct{}"
	}

	return strings.Join([]string{"CreateResWorkspaceRequest", string(data)}, " ")
}
