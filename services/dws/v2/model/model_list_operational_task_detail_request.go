package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOperationalTaskDetailRequest Request Object
type ListOperationalTaskDetailRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 分页偏移量，从0开始，页数减1。 **约束限制**： 不涉及。 **取值范围**： VacuumFull **默认取值**： VacuumFull
	Category *string `json:"category,omitempty"`

	// **参数解释**： 分页偏移量，从0开始，页数减1。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释**： 状态。 **约束限制**： 不涉及。 **取值范围**： running、waiting、finished、canceled **默认取值**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 开始日期。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 结束日期。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 分页查询，每页大小。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 100
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 分页偏移量，从0开始，页数减1。 **约束限制**： 不涉及。 **取值范围**： 大于等于0 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListOperationalTaskDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOperationalTaskDetailRequest struct{}"
	}

	return strings.Join([]string{"ListOperationalTaskDetailRequest", string(data)}, " ")
}
