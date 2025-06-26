package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadPlanStageReqWorkloadPlanStage **参数解释**： 资源管理计划阶段详情。 **取值范围**： 不涉及。
type WorkloadPlanStageReqWorkloadPlanStage struct {

	// **参数解释**： 日期。 **取值范围**： 不涉及。
	Day *string `json:"day,omitempty"`

	// **参数解释**： 月份。 **取值范围**： 不涉及。
	Month *string `json:"month,omitempty"`

	// **参数解释**： 计划阶段。 **取值范围**： 不涉及。
	StageName *string `json:"stage_name,omitempty"`

	// **参数解释**： 开始时间。 **取值范围**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 结束时间。 **取值范围**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 资源队列。 **取值范围**： 不涉及。
	QueueList *[]QueueResourceItem `json:"queue_list,omitempty"`
}

func (o WorkloadPlanStageReqWorkloadPlanStage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadPlanStageReqWorkloadPlanStage struct{}"
	}

	return strings.Join([]string{"WorkloadPlanStageReqWorkloadPlanStage", string(data)}, " ")
}
