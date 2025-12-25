package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePipeRequest Request Object
type UpdatePipeRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 管道ID
	PipeId string `json:"pipe_id"`

	Body *UpdatePipeRequestBody `json:"body,omitempty"`
}

func (o UpdatePipeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePipeRequest struct{}"
	}

	return strings.Join([]string{"UpdatePipeRequest", string(data)}, " ")
}
