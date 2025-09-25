package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecutionTimeDetailsInfo struct {
	ResourceTime *ResourceTime `json:"resource_time"`

	KernelTime *KernelTime `json:"kernel_time"`

	KernelExecutionTime *KernelExecutionTime `json:"kernel_execution_time,omitempty"`

	WaitEventTime *WaitEventTime `json:"wait_event_time"`
}

func (o ExecutionTimeDetailsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutionTimeDetailsInfo struct{}"
	}

	return strings.Join([]string{"ExecutionTimeDetailsInfo", string(data)}, " ")
}
