package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecutionResultList 函数流节点执行详细信息
type ExecutionResultList struct {

	// 流程节点ID。
	NodeId *string `json:"node_id,omitempty"`

	// 节点开始执行时间。
	BeginTime *int64 `json:"begin_time,omitempty"`

	// 节点执行结束时间。
	EndTime *int64 `json:"end_time,omitempty"`

	// FunctionGraph的执行id。
	FunctionExecutionId *string `json:"function_execution_id,omitempty"`

	// 节点输出。
	Output *interface{} `json:"output,omitempty"`

	// 节点状态。
	Status *string `json:"status,omitempty"`
}

func (o ExecutionResultList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutionResultList struct{}"
	}

	return strings.Join([]string{"ExecutionResultList", string(data)}, " ")
}
