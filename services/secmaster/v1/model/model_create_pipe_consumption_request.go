package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePipeConsumptionRequest Request Object
type CreatePipeConsumptionRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 管道ID
	PipeId string `json:"pipe_id"`

	Body *EnableConsumptionRequestBody `json:"body,omitempty"`
}

func (o CreatePipeConsumptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipeConsumptionRequest struct{}"
	}

	return strings.Join([]string{"CreatePipeConsumptionRequest", string(data)}, " ")
}
