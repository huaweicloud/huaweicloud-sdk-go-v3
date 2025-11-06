package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmActionRequestBody
type UpdateAlarmActionRequestBody struct {

	// 告警是否启用。true：启动。false：停止
	AlarmEnabled bool `json:"alarm_enabled"`
}

func (o UpdateAlarmActionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmActionRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAlarmActionRequestBody", string(data)}, " ")
}
