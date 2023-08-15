package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkflowInstanceRequest Request Object
type ShowWorkflowInstanceRequest struct {

	// 工作流实例名称。
	ExecutionName string `json:"execution_name"`

	// 工作流名称。
	GraphName string `json:"graph_name"`
}

func (o ShowWorkflowInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowInstanceRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkflowInstanceRequest", string(data)}, " ")
}
