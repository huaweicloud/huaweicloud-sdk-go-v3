package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ApplyWorkspaceRequest struct {
	Body *ApplyWorkspaceReq `json:"body,omitempty"`
}

func (o ApplyWorkspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyWorkspaceRequest struct{}"
	}

	return strings.Join([]string{"ApplyWorkspaceRequest", string(data)}, " ")
}
