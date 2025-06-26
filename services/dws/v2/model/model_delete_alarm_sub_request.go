package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlarmSubRequest Request Object
type DeleteAlarmSubRequest struct {

	// **参数解释**： 告警订阅ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	AlarmSubId string `json:"alarm_sub_id"`
}

func (o DeleteAlarmSubRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmSubRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlarmSubRequest", string(data)}, " ")
}
