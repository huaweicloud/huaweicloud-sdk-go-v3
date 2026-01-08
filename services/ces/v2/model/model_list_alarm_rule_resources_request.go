package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmRuleResourcesRequest Request Object
type ListAlarmRuleResourcesRequest struct {

	// **参数解释**： 告警规则ID。 **约束限制**： 不涉及。 **取值范围**： 以al开头，后跟22位的字母或数字。          **默认取值**： 不涉及。
	AlarmId string `json:"alarm_id"`

	// **参数解释**： 分页偏移量。 **约束限制**： 不涉及。 **取值范围**： 最小值为0，最大值为10000。 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分页大小。 **约束限制**： 不涉及。 **取值范围**： 最小值为1，最大值为100。 **默认取值**： 10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAlarmRuleResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmRuleResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmRuleResourcesRequest", string(data)}, " ")
}
