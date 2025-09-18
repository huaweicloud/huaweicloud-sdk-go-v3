package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineSchedule struct {

	// **参数解释**： 定时任务ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Uuid *string `json:"uuid,omitempty"`

	// **参数解释**： 任务类型。 **约束限制**： 不涉及。 **取值范围**： 只支持fixed。 **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 任务名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 是否启用。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Enable *bool `json:"enable,omitempty"`

	// **参数解释**： 一周内具体执行日。周日至周六对应1-7。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DaysOfWeek *[]int32 `json:"days_of_week,omitempty"`

	// **参数解释**： 时区。 **约束限制**： 不涉及。 **取值范围**： - \"China Standard Time\"。 - \"GMT Standard Time\"。 - \"South Africa Standard Time\"。 - \"Russian Standard Time\"。 - \"SE Asia Standard Time\"。  - \"Singapore Standard Time\"。 - \"Pacific SA Standard Time\"。 - \"E. South America Standard Time\"。  - \"Central Standard Time (Mexico)\"。 - \"Egypt Standard Time\"。 - \"Saudi Arabia Standard Time\"。 **默认取值**： 不涉及。
	TimeZone *string `json:"time_zone,omitempty"`
}

func (o PipelineSchedule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineSchedule struct{}"
	}

	return strings.Join([]string{"PipelineSchedule", string(data)}, " ")
}
