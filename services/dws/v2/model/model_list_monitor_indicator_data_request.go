package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMonitorIndicatorDataRequest Request Object
type ListMonitorIndicatorDataRequest struct {

	// **参数解释**： 开始时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	From string `json:"from"`

	// **参数解释**： 结束时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	To string `json:"to"`

	// **参数解释**： 取值方法。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Function *string `json:"function,omitempty"`

	// **参数解释**： 取值周期。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Period *string `json:"period,omitempty"`

	// **参数解释**： 指标名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	IndicatorName string `json:"indicator_name"`

	// **参数解释**： 第一层级。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Dim0 string `json:"dim0"`

	// **参数解释**： 第二层级。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Dim1 *string `json:"dim1,omitempty"`
}

func (o ListMonitorIndicatorDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonitorIndicatorDataRequest struct{}"
	}

	return strings.Join([]string{"ListMonitorIndicatorDataRequest", string(data)}, " ")
}
