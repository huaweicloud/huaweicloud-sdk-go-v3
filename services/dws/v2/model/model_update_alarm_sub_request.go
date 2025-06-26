package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmSubRequest Request Object
type UpdateAlarmSubRequest struct {

	// **参数解释**： 告警订阅ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	AlarmSubId string `json:"alarm_sub_id"`

	Body *AlarmSubUpdateRequest `json:"body,omitempty"`
}

func (o UpdateAlarmSubRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmSubRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlarmSubRequest", string(data)}, " ")
}
