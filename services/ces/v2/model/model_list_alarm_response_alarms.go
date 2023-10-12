package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAlarmResponseAlarms struct {

	// 告警规则id，以al开头，包含22个数字或字母
	AlarmId *string `json:"alarm_id,omitempty"`

	// 告警名称, 只能包含0-9/a-z/A-Z/_/-或汉字，长度1-128
	Name *string `json:"name,omitempty"`

	// 告警描述，长度0-256
	Description *string `json:"description,omitempty"`

	// 查询服务的命名空间，各服务命名空间请参考[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)
	Namespace *string `json:"namespace,omitempty"`

	// 告警策略
	Policies *[]Policy `json:"policies,omitempty"`

	// 资源列表，关联资源需要使用查询告警规则资源接口获取
	Resources *[]ResourcesInListResp `json:"resources,omitempty"`

	Type *AlarmType `json:"type,omitempty"`

	// 告警开关
	Enabled *bool `json:"enabled,omitempty"`

	// 是否开启告警通知
	NotificationEnabled *bool `json:"notification_enabled,omitempty"`

	// 告警触发的动作
	AlarmNotifications *[]Notification `json:"alarm_notifications,omitempty"`

	// 告警恢复触发的动作
	OkNotifications *[]Notification `json:"ok_notifications,omitempty"`

	// 告警通知开启时间
	NotificationBeginTime *string `json:"notification_begin_time,omitempty"`

	// 告警通知关闭时间
	NotificationEndTime *string `json:"notification_end_time,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 告警规则关联告警模板ID
	AlarmTemplateId *string `json:"alarm_template_id,omitempty"`
}

func (o ListAlarmResponseAlarms) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmResponseAlarms struct{}"
	}

	return strings.Join([]string{"ListAlarmResponseAlarms", string(data)}, " ")
}
