package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HandleAlarmRequest Request Object
type HandleAlarmRequest struct {

	// 告警ID
	AlarmId string `json:"alarm_id"`

	Body *AutoHandlerAlarmRequestBody `json:"body,omitempty"`
}

func (o HandleAlarmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleAlarmRequest struct{}"
	}

	return strings.Join([]string{"HandleAlarmRequest", string(data)}, " ")
}
