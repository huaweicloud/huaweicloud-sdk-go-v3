package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePipeConsumptionRequest Request Object
type DeletePipeConsumptionRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 管道ID
	PipeId string `json:"pipe_id"`
}

func (o DeletePipeConsumptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePipeConsumptionRequest struct{}"
	}

	return strings.Join([]string{"DeletePipeConsumptionRequest", string(data)}, " ")
}
