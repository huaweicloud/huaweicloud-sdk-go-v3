package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlarmRequest Request Object
type ShowAlarmRequest struct {

	// **参数解释**： 告警规则ID， **约束限制**： 不涉及。 **取值范围**： 以al开头，后跟22位由字母或数字组成的字符串。长度为24个字符。 **默认取值**： 不涉及。
	AlarmId string `json:"alarm_id"`
}

func (o ShowAlarmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmRequest struct{}"
	}

	return strings.Join([]string{"ShowAlarmRequest", string(data)}, " ")
}
