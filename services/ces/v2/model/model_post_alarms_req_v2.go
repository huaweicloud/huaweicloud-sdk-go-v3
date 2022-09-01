package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PostAlarmsReqV2 struct {

	// 告警名称, 只能包含0-9/a-z/A-Z/_/-或汉字，长度1-128
	Name string `json:"name" xml:"name"`

	// 告警描述，长度0-256
	Description *string `json:"description,omitempty" xml:"description"`

	// 查询服务的命名空间，各服务命名空间请参考[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)
	Namespace string `json:"namespace" xml:"namespace"`

	// 资源分组ID，监控范围为资源分组时必传
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id"`

	// 资源列表，监控范围为指定资源时必传
	Resources [][]Dimension `json:"resources" xml:"resources"`

	// 告警策略
	Policies []Policy `json:"policies" xml:"policies"`

	// 告警规则类型
	Type string `json:"type" xml:"type"`

	// 告警触发的动作
	AlarmNotifications *[]Notification `json:"alarm_notifications,omitempty" xml:"alarm_notifications"`

	// 告警恢复触发的动作
	OkNotifications *[]Notification `json:"ok_notifications,omitempty" xml:"ok_notifications"`

	// 告警通知开启时间
	NotificationBeginTime *string `json:"notification_begin_time,omitempty" xml:"notification_begin_time"`

	// 告警通知关闭时间
	NotificationEndTime *string `json:"notification_end_time,omitempty" xml:"notification_end_time"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 告警开关
	Enabled bool `json:"enabled" xml:"enabled"`

	// 是否开启告警通知
	NotificationEnabled bool `json:"notification_enabled" xml:"notification_enabled"`

	// 告警规则关联告警模板ID，如果传了，告警规则关联的策略会和告警模板策略联动变化
	AlarmTemplateId *string `json:"alarm_template_id,omitempty" xml:"alarm_template_id"`
}

func (o PostAlarmsReqV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostAlarmsReqV2 struct{}"
	}

	return strings.Join([]string{"PostAlarmsReqV2", string(data)}, " ")
}