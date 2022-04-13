package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ModifyAlarmActionReq struct {
	// 告警是否启用。true：启动。false：停止

	AlarmEnabled bool `json:"alarm_enabled"`
}

func (o ModifyAlarmActionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyAlarmActionReq struct{}"
	}

	return strings.Join([]string{"ModifyAlarmActionReq", string(data)}, " ")
}
