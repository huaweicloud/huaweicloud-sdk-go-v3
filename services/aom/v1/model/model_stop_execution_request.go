package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopExecutionRequest Request Object
type StopExecutionRequest struct {

	// 任务id，从工作流命令列表中获取的工作流id。
	WorkflowId string `json:"workflow_id"`

	// 任务执行id。
	ExecutionId string `json:"execution_id"`
}

func (o StopExecutionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopExecutionRequest struct{}"
	}

	return strings.Join([]string{"StopExecutionRequest", string(data)}, " ")
}
