package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PostAlarmsReqV2 struct {

	// 告警名称, 只能包含0-9/a-z/A-Z/_/-或汉字，长度1-128
	Name string `json:"name"`

	// 告警描述，长度0-256
	Description *string `json:"description,omitempty"`

	// 查询服务的命名空间，各服务命名空间请参考[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)
	Namespace string `json:"namespace"`

	// 资源分组ID，以rg开头，后跟22位由字母或数字组成的字符串
	ResourceGroupId *string `json:"resource_group_id,omitempty"`

	// 资源列表，监控范围为指定资源时必传
	Resources [][]Dimension `json:"resources"`

	// 告警规则关联告警模板ID
	AlarmTemplateId *string `json:"alarm_template_id,omitempty"`

	// 租户标签列表
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// 告警策略，当alarm_template_id字段为空时必填，不为空时不填
	Policies *[]Policy `json:"policies,omitempty"`

	Type *AlarmType `json:"type"`

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

	// 告警开关
	Enabled bool `json:"enabled"`

	// 是否开启告警通知
	NotificationEnabled bool `json:"notification_enabled"`
}

func (o PostAlarmsReqV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostAlarmsReqV2 struct{}"
	}

	return strings.Join([]string{"PostAlarmsReqV2", string(data)}, " ")
}
