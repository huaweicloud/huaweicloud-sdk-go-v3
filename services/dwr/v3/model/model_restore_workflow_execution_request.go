package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreWorkflowExecutionRequest Request Object
type RestoreWorkflowExecutionRequest struct {

	// 工作流实例名。
	ExecutionName string `json:"execution_name"`

	// 工作流名。
	GraphName string `json:"graph_name"`
}

func (o RestoreWorkflowExecutionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreWorkflowExecutionRequest struct{}"
	}

	return strings.Join([]string{"RestoreWorkflowExecutionRequest", string(data)}, " ")
}
