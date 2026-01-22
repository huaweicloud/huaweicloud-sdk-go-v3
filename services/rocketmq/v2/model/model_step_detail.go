package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StepDetail struct {

	// **参数解释**： 任务名称。  **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 任务状态。 **约束限制**： 不涉及。 **取值范围**： - COMPLETED：任务已完成。 - IN_PROGRESS：任务正在进行。 - FAILED：任务失败。 - WAITING：等待开始。 **默认取值**： 不涉及。
	Statue *string `json:"statue,omitempty"`

	// **参数解释**： 开始时间。    **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 结束时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`
}

func (o StepDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StepDetail struct{}"
	}

	return strings.Join([]string{"StepDetail", string(data)}, " ")
}
