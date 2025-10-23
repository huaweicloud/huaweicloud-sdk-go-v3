package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmRuleEntity 告警规则实体类
type AlarmRuleEntity struct {

	// 告警规则ID
	AlarmId string `json:"alarm_id"`

	// 告警名称
	Name *string `json:"name,omitempty"`

	// 告警描述
	Description *string `json:"description,omitempty"`

	Namespace *AlarmNamespaceEnum `json:"namespace,omitempty"`

	// 告警策略
	Policies *string `json:"policies,omitempty"`

	// 资源列表
	Resources *string `json:"resources,omitempty"`

	// 告警类型
	Type *string `json:"type,omitempty"`

	// 告警开关
	Enabled *bool `json:"enabled,omitempty"`

	// 是否开启告警通知
	NotificationEnabled *bool `json:"notification_enabled,omitempty"`

	// 告警触发的动作
	AlarmNotifications *string `json:"alarm_notifications,omitempty"`

	// 告警恢复触发的动作
	OkNotifications *string `json:"ok_notifications,omitempty"`

	// 告警通知开启时间
	NotificationBeginTime *string `json:"notification_begin_time,omitempty"`

	// 告警通知关闭时间
	NotificationEndTime *string `json:"notification_end_time,omitempty"`

	// RegionID
	RegionId *string `json:"region_id,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 告警规则关联告警模板ID
	AlarmTemplateId *string `json:"alarm_template_id,omitempty"`
}

func (o AlarmRuleEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmRuleEntity struct{}"
	}

	return strings.Join([]string{"AlarmRuleEntity", string(data)}, " ")
}
