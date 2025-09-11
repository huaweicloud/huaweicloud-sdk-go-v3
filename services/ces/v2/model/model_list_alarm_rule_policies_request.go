package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmRulePoliciesRequest Request Object
type ListAlarmRulePoliciesRequest struct {

	// **参数解释**： 告警规则ID。     **约束限制**： 不涉及。 **取值范围**： 以al开头，后跟22个数字或字母。字符长度为24 **默认取值**： 不涉及。
	AlarmId string `json:"alarm_id"`

	// **参数解释**： 分页偏移量。 **约束限制**： 不涉及。 **取值范围**： 0-10000 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分页大小。 **约束限制**： 不涉及。 **取值范围**： 1-100 **默认取值**： 10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAlarmRulePoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmRulePoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmRulePoliciesRequest", string(data)}, " ")
}
