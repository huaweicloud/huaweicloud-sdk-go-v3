package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KernelTimeDetails 内核耗时详细信息
type KernelTimeDetails struct {

	// **参数解释**: SQL解析时间（单位：微秒）。 **取值范围**: 不涉及。
	ParseTime int64 `json:"parse_time"`

	// **参数解释**: SQL重写时间（单位：微秒）。 **取值范围**: 不涉及。
	RewriteTime int64 `json:"rewrite_time"`

	// **参数解释**: SQL生成计划时间（单位：微秒）。 **取值范围**: 不涉及。
	PlanTime int64 `json:"plan_time"`

	// **参数解释**: 执行器内执行时间（单位：微秒）。 **取值范围**: 不涉及。
	ExecutionTime int64 `json:"execution_time"`

	// **参数解释**: 其余耗时（单位：微秒）。 **取值范围**: 不涉及。
	OtherTime int64 `json:"other_time"`
}

func (o KernelTimeDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KernelTimeDetails struct{}"
	}

	return strings.Join([]string{"KernelTimeDetails", string(data)}, " ")
}
