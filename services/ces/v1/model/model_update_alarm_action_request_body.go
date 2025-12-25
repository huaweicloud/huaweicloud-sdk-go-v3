package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmActionRequestBody
type UpdateAlarmActionRequestBody struct {

	// **参数解释**： 告警是否启用 **约束限制**： 不涉及 **取值范围**： - true：开启告警 - false：关闭告警 **默认取值**： 不涉及
	AlarmEnabled bool `json:"alarm_enabled"`
}

func (o UpdateAlarmActionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmActionRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAlarmActionRequestBody", string(data)}, " ")
}
