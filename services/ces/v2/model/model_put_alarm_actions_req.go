package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PutAlarmActionsReq struct {
	// 告警名称, 只能包含0-9/a-z/A-Z/_/-或汉字，长度1-128

	Name string `json:"name"`
	// 告警描述，长度0-256

	Description *string `json:"description,omitempty"`
	// 是否开启告警通知

	ActionEnabled *string `json:"action_enabled,omitempty"`
	// 告警触发的动作

	AlarmActions *[]SmnAction `json:"alarm_actions,omitempty"`
	// 告警恢复触发的动作

	OkActions *[]SmnAction `json:"ok_actions,omitempty"`
	// 告警通知开启时间

	ActionBeginTime *string `json:"action_begin_time,omitempty"`
	// 告警通知关闭时间

	ActionEndTime *string `json:"action_end_time,omitempty"`
}

func (o PutAlarmActionsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutAlarmActionsReq struct{}"
	}

	return strings.Join([]string{"PutAlarmActionsReq", string(data)}, " ")
}
