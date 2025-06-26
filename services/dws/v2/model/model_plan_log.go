package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlanLog **参数解释**： 工作计划日志。 **取值范围**： 不涉及。
type PlanLog struct {

	// **参数解释**： 执行时间。 **取值范围**： 不涉及。
	ExecTime string `json:"exec_time"`

	// **参数解释**： 执行计划阶段。 **取值范围**： 不涉及。
	StageInfo string `json:"stage_info"`

	// **参数解释**： 执行结果。 **取值范围**： 不涉及。
	ExecResult int32 `json:"exec_result"`

	// **参数解释**： 执行日志。 **取值范围**： 不涉及。
	ExecLog string `json:"exec_log"`
}

func (o PlanLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlanLog struct{}"
	}

	return strings.Join([]string{"PlanLog", string(data)}, " ")
}
