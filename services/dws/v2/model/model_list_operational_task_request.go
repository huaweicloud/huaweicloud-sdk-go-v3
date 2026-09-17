package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOperationalTaskRequest Request Object
type ListOperationalTaskRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 时区偏移量，默认为UTC时间（0偏移），即+0000。 **约束限制**： 不涉及。 **取值范围**： -2359~+2359 **默认取值**： +0000
	TimeZone *string `json:"time_zone,omitempty"`

	// **参数解释**： 任务类型。 **约束限制**： 仅支持两种值。 **取值范围**： Date：单次型任务； Window：周期型任务； **默认取值**： 无。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 分页查询，每页大小。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 100
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 分页偏移量，从0开始，页数减1。 **约束限制**： 不涉及。 **取值范围**： 大于等于0 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListOperationalTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOperationalTaskRequest struct{}"
	}

	return strings.Join([]string{"ListOperationalTaskRequest", string(data)}, " ")
}
