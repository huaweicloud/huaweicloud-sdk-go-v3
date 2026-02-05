package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePipeIndexRequest Request Object
type UpdatePipeIndexRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 管道ID
	PipeId string `json:"pipe_id"`

	Body *UpdateIndexRequestBody `json:"body,omitempty"`
}

func (o UpdatePipeIndexRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePipeIndexRequest struct{}"
	}

	return strings.Join([]string{"UpdatePipeIndexRequest", string(data)}, " ")
}
