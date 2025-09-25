package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KernelExecTimeDetails 内核耗时详细信息
type KernelExecTimeDetails struct {

	// **参数解释**: 执行器内执行时间（单位：微秒）。 **取值范围**: 不涉及。
	ExecutionTime int64 `json:"execution_time"`

	// **参数解释**: 执行器内执行时间（单位：微秒）。 **取值范围**: 不涉及。
	OtherTime int64 `json:"other_time"`
}

func (o KernelExecTimeDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KernelExecTimeDetails struct{}"
	}

	return strings.Join([]string{"KernelExecTimeDetails", string(data)}, " ")
}
