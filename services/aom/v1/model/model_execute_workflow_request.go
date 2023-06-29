package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteWorkflowRequest Request Object
type ExecuteWorkflowRequest struct {

	// 任务id，从工作流命令列表中获取的工作流id。
	WorkflowId string `json:"workflow_id"`
}

func (o ExecuteWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteWorkflowRequest struct{}"
	}

	return strings.Join([]string{"ExecuteWorkflowRequest", string(data)}, " ")
}
