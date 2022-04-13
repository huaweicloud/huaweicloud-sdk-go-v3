package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowWorkFlowRequest struct {
	// 函数工作流ID

	WorkflowId string `json:"workflow_id"`
}

func (o ShowWorkFlowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkFlowRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkFlowRequest", string(data)}, " ")
}
