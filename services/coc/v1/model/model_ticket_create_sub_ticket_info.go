package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TicketCreateSubTicketInfo struct {

	// **参数解释：** 变更服务。 **约束限制：** 当target_type为change_scope时，该字段需传递对应的变更服务中文名称。 **取值范围：** 不涉及 **默认取值：** 不涉及
	AppName *string `json:"app_name,omitempty"`

	// **参数解释：** 变更Region。 **约束限制：** 当target_type为change_scope时，该字段需传递对应的变更RegionID。 **取值范围：** 不涉及 **默认取值：** 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释：** 目标选项信息，该值可传递不同的信息，当传递变更应用时，该值传递change_scope，当传递变更计划时，该值传递child_ticket。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TargetType *string `json:"target_type,omitempty"`

	// **参数解释：** 传递变更单RegionID，需和target_type配合使用。 **约束限制：** 当target_type值为child_ticket时，该值有效。 **取值范围：** 不涉及 **默认取值：** 不涉及
	TargetValue *string `json:"target_value,omitempty"`

	// **参数解释：** 变更单计划开始时间时间戳。 **约束限制：** 当target_type值为child_ticket时，该值有效。 **取值范围：** 不涉及 **默认取值：** 不涉及
	ExpectedStartTime *int64 `json:"expected_start_time,omitempty"`

	// **参数解释：** 变更单计划结束时间时间戳。 **约束限制：** 当target_type值为child_ticket时，该值有效。 **取值范围：** 不涉及 **默认取值：** 不涉及
	ExpectedEndTime *int64 `json:"expected_end_time,omitempty"`

	// **参数解释：** 变更单实施人。 **约束限制：** 当target_type值为child_ticket时，该值有效。 **取值范围：** 不涉及 **默认取值：** 不涉及
	Executors *string `json:"executors,omitempty"`

	// **参数解释：** 变更单配合人。 **约束限制：** 当target_type值为child_ticket时，该值有效。 **取值范围：** 不涉及 **默认取值：** 不涉及
	Cooperators *string `json:"cooperators,omitempty"`
}

func (o TicketCreateSubTicketInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TicketCreateSubTicketInfo struct{}"
	}

	return strings.Join([]string{"TicketCreateSubTicketInfo", string(data)}, " ")
}
