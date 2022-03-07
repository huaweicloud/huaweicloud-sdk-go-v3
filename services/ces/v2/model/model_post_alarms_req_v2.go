package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PostAlarmsReqV2 struct {
	// 告警名称, 只能包含0-9/a-z/A-Z/_/-或汉字，长度1-128

	Name string `json:"name"`
	// 告警规则描述

	Description *string `json:"description,omitempty"`
	// 命名空间

	Namespace string `json:"namespace"`
	// 资源分组ID

	ResourceGroupId *string `json:"resource_group_id,omitempty"`
	// 资源信息

	Resources [][]Dimension `json:"resources"`
	// 策略信息

	Policies []Policy `json:"policies"`
	// 告警规则类型

	Type *string `json:"type,omitempty"`
	// 告警通知

	AlarmActions *[]SmnAction `json:"alarm_actions,omitempty"`
	// 告警恢复通知

	OkActions *[]SmnAction `json:"ok_actions,omitempty"`
	// 告警通知开始时间

	ActionBeginTime *string `json:"action_begin_time,omitempty"`
	// 告警通知结束时间

	ActionEndTime *string `json:"action_end_time,omitempty"`
	// 企业项目ID

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 告警开关

	Enabled bool `json:"enabled"`
	// 告警通知开关

	ActionEnabled bool `json:"action_enabled"`
}

func (o PostAlarmsReqV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostAlarmsReqV2 struct{}"
	}

	return strings.Join([]string{"PostAlarmsReqV2", string(data)}, " ")
}
