package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlarmRequest Request Object
type DeleteAlarmRequest struct {

	// 告警规则的ID。
	AlarmId string `json:"alarm_id"`
}

func (o DeleteAlarmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlarmRequest", string(data)}, " ")
}
