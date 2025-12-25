package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateOneClickAlarmsEnabledStateRequestBody struct {

	// **参数解释** 需要批量启停的告警规则的ID列表 **约束限制** 包含的告警规则ID个数为1到100
	AlarmIds []string `json:"alarm_ids"`

	// **参数解释**： 是否开启告警规则。     **约束限制**： 不涉及。 **取值范围**： 布尔值。 - true:开启。 - false:关闭。 **默认取值**： true
	AlarmEnabled bool `json:"alarm_enabled"`

	// **参数解释** 一键告警中的规则全部被停用时是否保留规则信息 **约束限制** 不涉及 **取值范围** - true:保留规则信息 - false:不保留规则信息 **默认取值** false
	RetainWhenAllDisabled *bool `json:"retain_when_all_disabled,omitempty"`

	// **参数解释** 当传此参数时，将会批量停用一键告警规则 **约束限制** 不涉及 **取值范围** - disable: 表示将会批量停用一键告警规则 **默认取值** 不涉及
	ActionType *string `json:"action_type,omitempty"`
}

func (o BatchUpdateOneClickAlarmsEnabledStateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateOneClickAlarmsEnabledStateRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateOneClickAlarmsEnabledStateRequestBody", string(data)}, " ")
}
