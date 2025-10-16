package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateOneClickAlarmsEnabledStateRequestBody struct {

	// 需要批量启停的告警规则的ID列表
	AlarmIds []string `json:"alarm_ids"`

	// **参数解释**： 是否开启告警规则。     **约束限制**： 不涉及。 **取值范围**： 布尔值。 - true:开启。 - false:关闭。 **默认取值**： true
	AlarmEnabled bool `json:"alarm_enabled"`

	// 一键告警中的规则全部被停用时是否保留规则信息。true:保留；false:删除。
	RetainWhenAllDisabled *bool `json:"retain_when_all_disabled,omitempty"`
}

func (o BatchUpdateOneClickAlarmsEnabledStateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateOneClickAlarmsEnabledStateRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateOneClickAlarmsEnabledStateRequestBody", string(data)}, " ")
}
