package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmRulePoliciesRequest Request Object
type UpdateAlarmRulePoliciesRequest struct {

	// **参数解释**： 告警规则ID。 **约束限制**： 不涉及。 **取值范围**： 以al开头，只能由大写字母、小写字母、数字组成，且长度为24个字符。           **默认取值**： 不涉及。
	AlarmId string `json:"alarm_id"`

	Body *UpdateAlarmRulePoliciesReqBodyV2 `json:"body,omitempty"`
}

func (o UpdateAlarmRulePoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmRulePoliciesRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlarmRulePoliciesRequest", string(data)}, " ")
}
