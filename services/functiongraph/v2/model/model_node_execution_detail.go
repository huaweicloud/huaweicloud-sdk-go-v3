package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 函数流节点执行详细信息
type NodeExecutionDetail struct {

	// 流程节点ID
	NodeId *string `json:"node_id,omitempty" xml:"node_id"`

	// 流程节点名称
	NodeName *string `json:"node_name,omitempty" xml:"node_name"`

	// 流程节点执行ID
	ExecutionId *string `json:"execution_id,omitempty" xml:"execution_id"`

	// 节点执行记录
	Executions *[]NodeExecution `json:"executions,omitempty" xml:"executions"`
}

func (o NodeExecutionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeExecutionDetail struct{}"
	}

	return strings.Join([]string{"NodeExecutionDetail", string(data)}, " ")
}
