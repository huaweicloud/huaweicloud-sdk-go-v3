package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 函数流节点执行详细信息
type NodeExecutionDetail struct {
	// 流程节点ID

	NodeId *string `json:"node_id,omitempty"`
	// 节点执行记录

	Executions *[]NodeExecution `json:"executions,omitempty"`
}

func (o NodeExecutionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeExecutionDetail struct{}"
	}

	return strings.Join([]string{"NodeExecutionDetail", string(data)}, " ")
}
