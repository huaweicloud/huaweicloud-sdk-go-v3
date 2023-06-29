package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchWorkflowExecutionDetailRequest Request Object
type SearchWorkflowExecutionDetailRequest struct {

	// 工作流ID，唯一标识，根据project_id和workflow_name生成。
	WorkflowId string `json:"workflow_id"`

	// 工作流执行ID。
	ExecutionId string `json:"execution_id"`
}

func (o SearchWorkflowExecutionDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchWorkflowExecutionDetailRequest struct{}"
	}

	return strings.Join([]string{"SearchWorkflowExecutionDetailRequest", string(data)}, " ")
}
