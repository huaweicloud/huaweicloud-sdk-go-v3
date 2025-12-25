package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateOneClickAlarmPoliciesEnabledStateRequest Request Object
type BatchUpdateOneClickAlarmPoliciesEnabledStateRequest struct {

	// **参数解释** 一键告警ID **约束限制** 不涉及 **取值范围** 长度为1到64字符，只能包含字母数字 **默认取值** 不涉及
	OneClickAlarmId string `json:"one_click_alarm_id"`

	// **参数解释** 告警规则ID **约束限制** 不涉及 **取值范围** 以al开头，后跟22个数字或字母。 **默认取值** 不涉及
	AlarmId string `json:"alarm_id"`

	Body *BatchEnableAlarmPoliciesRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateOneClickAlarmPoliciesEnabledStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateOneClickAlarmPoliciesEnabledStateRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateOneClickAlarmPoliciesEnabledStateRequest", string(data)}, " ")
}
