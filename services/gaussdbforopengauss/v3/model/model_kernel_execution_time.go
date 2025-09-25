package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KernelExecutionTime 内核执行模块耗时
type KernelExecutionTime struct {

	// **参数解释**: 总耗时（单位：微秒）。 **取值范围**: 不涉及。
	AllTime int64 `json:"all_time"`

	KernelExecutionTimeDetails *KernelExecTimeDetails `json:"kernel_execution_time_details"`
}

func (o KernelExecutionTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KernelExecutionTime struct{}"
	}

	return strings.Join([]string{"KernelExecutionTime", string(data)}, " ")
}
