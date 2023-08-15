package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmSubRequest Request Object
type UpdateAlarmSubRequest struct {

	// 告警订阅ID
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
