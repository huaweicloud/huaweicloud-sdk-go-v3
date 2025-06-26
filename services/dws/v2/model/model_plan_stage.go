package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlanStage **参数解释**： 工作计划阶段。 **取值范围**： 不涉及。
type PlanStage struct {

	// **参数解释**： 计划月份。 **取值范围**： 不涉及。
	Month string `json:"month"`

	// **参数解释**： 计划日期。 **取值范围**： 不涉及。
	Day string `json:"day"`

	// **参数解释**： 计划ID。 **取值范围**： 不涉及。
	PlanId string `json:"plan_id"`

	// **参数解释**： 计划阶段ID。 **取值范围**： 不涉及。
	StageId string `json:"stage_id"`

	// **参数解释**： 计划阶段名称。 **取值范围**： 不涉及。
	StageName string `json:"stage_name"`

	// **参数解释**： 计划开始时间。 **取值范围**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 计划结束时间。 **取值范围**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 下次校验时间。 **取值范围**： 不涉及。
	NextValidTime *string `json:"next_valid_time,omitempty"`

	// **参数解释**： 资源队列列表。 **取值范围**： 不涉及。
	QueueList *[]QueueResourceItem `json:"queue_list,omitempty"`
}

func (o PlanStage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlanStage struct{}"
	}

	return strings.Join([]string{"PlanStage", string(data)}, " ")
}
