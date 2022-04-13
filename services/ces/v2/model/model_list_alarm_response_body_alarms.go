package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAlarmResponseBodyAlarms struct {
	// 告警规则ID

	AlarmId *string `json:"alarm_id,omitempty"`
	// 告警规则名称

	Name *string `json:"name,omitempty"`
	// 告警规则描述

	Description *string `json:"description,omitempty"`
	// 告警规则的命名空间

	Namespace *string `json:"namespace,omitempty"`
	// 策略

	Policies *[]Policy `json:"policies,omitempty"`
	// 资源

	Resources *[]ResourcesInListResp `json:"resources,omitempty"`
	// 告警类型

	Type *string `json:"type,omitempty"`
	// 告警开关

	Enabled *bool `json:"enabled,omitempty"`
	// 告警通知开关

	ActionEnabled *bool `json:"action_enabled,omitempty"`
	// 告警通知

	AlarmActions *[]SmnAction `json:"alarm_actions,omitempty"`
	// 告警恢复通知

	OkActions *[]SmnAction `json:"ok_actions,omitempty"`
	// 数据不足通知

	InsufficientdataActions *[]SmnAction `json:"insufficientdata_actions,omitempty"`
	// 告警通知起始时间

	ActionBeginTime *string `json:"action_begin_time,omitempty"`
	// 告警通知结束时间

	ActionEndTime *string `json:"action_end_time,omitempty"`
	// 告警更新时间，参考API规范，使用UTC时间

	UpdateTime *string `json:"update_time,omitempty"`
	// 一键告警标志

	OneClickAlarmFlag *int32 `json:"one_click_alarm_flag,omitempty"`
	// 企业项目ID

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListAlarmResponseBodyAlarms) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmResponseBodyAlarms struct{}"
	}

	return strings.Join([]string{"ListAlarmResponseBodyAlarms", string(data)}, " ")
}
