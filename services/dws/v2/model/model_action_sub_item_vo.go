package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionSubItemVo **参数解释**： 任务详情子项详情。 **取值范围**： 不涉及。
type ActionSubItemVo struct {

	// **参数解释**： 状态。 **取值范围**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 扩展信息。 **取值范围**： 不涉及。
	Detail *string `json:"detail,omitempty"`

	// **参数解释**： 子项名称，根据请求header中的x-language字段返回对应的中/英文名称。 **取值范围**： 不涉及。
	SubItemName *string `json:"sub_item_name,omitempty"`

	// **参数解释**： 开始时间。 **取值范围**： 时间格式，或null值。
	BeginTime *string `json:"begin_time,omitempty"`

	// **参数解释**： 结束时间。 **取值范围**： 时间格式，或null值。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 预估时间。 **取值范围**： 整数。
	EstimatedTime *int32 `json:"estimated_time,omitempty"`
}

func (o ActionSubItemVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionSubItemVo struct{}"
	}

	return strings.Join([]string{"ActionSubItemVo", string(data)}, " ")
}
