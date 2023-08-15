package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AsyncInvokeApiStartWorkflowRequest Request Object
type AsyncInvokeApiStartWorkflowRequest struct {

	// 工作流名称。
	GraphName string `json:"graph_name"`

	Body *ExecuteWorkflowBody `json:"body,omitempty"`
}

func (o AsyncInvokeApiStartWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AsyncInvokeApiStartWorkflowRequest struct{}"
	}

	return strings.Join([]string{"AsyncInvokeApiStartWorkflowRequest", string(data)}, " ")
}
