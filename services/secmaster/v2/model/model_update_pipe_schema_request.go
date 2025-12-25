package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePipeSchemaRequest Request Object
type UpdatePipeSchemaRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 管道ID
	PipeId string `json:"pipe_id"`

	Body *PipeSchema `json:"body,omitempty"`
}

func (o UpdatePipeSchemaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePipeSchemaRequest struct{}"
	}

	return strings.Join([]string{"UpdatePipeSchemaRequest", string(data)}, " ")
}
