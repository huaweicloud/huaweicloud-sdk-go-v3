package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipeConsumptionRequest Request Object
type ShowPipeConsumptionRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 管道ID
	PipeId string `json:"pipe_id"`
}

func (o ShowPipeConsumptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipeConsumptionRequest struct{}"
	}

	return strings.Join([]string{"ShowPipeConsumptionRequest", string(data)}, " ")
}
