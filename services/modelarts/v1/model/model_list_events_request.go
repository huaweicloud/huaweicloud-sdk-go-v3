package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventsRequest Request Object
type ListEventsRequest struct {

	// **参数解释**：事件所属资源类型。可选值为pools，表示资源池。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Resource string `json:"resource"`

	// **参数解释**：事件所属资源名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：单页查询最大数量，该值为空或者0时默认返回500条记录，单页最大允许查询500条记录。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：分页查询的上一页标记，内容为UUID字符串，查询第一页时为空。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Continue *string `json:"continue,omitempty"`

	// **参数解释**：事件开始时间戳。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Since *int32 `json:"since,omitempty"`

	// **参数解释**：事件结束时间戳。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Until *int32 `json:"until,omitempty"`

	// **参数解释**：事件类型。 **约束限制**：不涉及。 **取值范围**：可选值如下： - Normal：正常 - Warning：异常 **默认取值**：不涉及。
	Type *string `json:"type,omitempty"`
}

func (o ListEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventsRequest struct{}"
	}

	return strings.Join([]string{"ListEventsRequest", string(data)}, " ")
}
