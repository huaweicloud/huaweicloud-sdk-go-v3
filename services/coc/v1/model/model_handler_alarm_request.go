package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HandlerAlarmRequest Request Object
type HandlerAlarmRequest struct {

	// 告警ID
	AlarmId string `json:"alarm_id"`

	Body *AutoHandlerAlarmRequestBody `json:"body,omitempty"`
}

func (o HandlerAlarmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandlerAlarmRequest struct{}"
	}

	return strings.Join([]string{"HandlerAlarmRequest", string(data)}, " ")
}
