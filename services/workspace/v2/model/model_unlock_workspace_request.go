package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnlockWorkspaceRequest Request Object
type UnlockWorkspaceRequest struct {
	Body *UnlockWorkspaceRequestBody `json:"body,omitempty"`
}

func (o UnlockWorkspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnlockWorkspaceRequest struct{}"
	}

	return strings.Join([]string{"UnlockWorkspaceRequest", string(data)}, " ")
}
