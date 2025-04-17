package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListAlarmsRespAlarms struct {

	// 告警规则id，以al开头，包含22个数字或字母
	AlarmId *string `json:"alarm_id,omitempty"`

	// 告警名称, 只能包含0-9/a-z/A-Z/_/-或汉字，长度1-128
	Name *string `json:"name,omitempty"`

	// 告警描述，长度0-256
	Description *string `json:"description,omitempty"`

	// 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”
	Namespace *string `json:"namespace,omitempty"`

	// 告警策略
	Policies *[]OneClickAlarmPolicy `json:"policies,omitempty"`

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

	// NOTIFICATION_GROUP(通知组)/TOPIC_SUBSCRIPTION(主题订阅)/NOTIFICATION_POLICY(通知策略)
	NotificationManner *ListAlarmsRespAlarmsNotificationManner `json:"notification_manner,omitempty"`

	// 关联的通知策略ID列表
	NotificationPolicyIds *[]string `json:"notification_policy_ids,omitempty"`
}

func (o ListAlarmsRespAlarms) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmsRespAlarms struct{}"
	}

	return strings.Join([]string{"ListAlarmsRespAlarms", string(data)}, " ")
}

type ListAlarmsRespAlarmsNotificationManner struct {
	value string
}

type ListAlarmsRespAlarmsNotificationMannerEnum struct {
	NOTIFICATION_GROUP  ListAlarmsRespAlarmsNotificationManner
	TOPIC_SUBSCRIPTION  ListAlarmsRespAlarmsNotificationManner
	NOTIFICATION_POLICY ListAlarmsRespAlarmsNotificationManner
}

func GetListAlarmsRespAlarmsNotificationMannerEnum() ListAlarmsRespAlarmsNotificationMannerEnum {
	return ListAlarmsRespAlarmsNotificationMannerEnum{
		NOTIFICATION_GROUP: ListAlarmsRespAlarmsNotificationManner{
			value: "NOTIFICATION_GROUP",
		},
		TOPIC_SUBSCRIPTION: ListAlarmsRespAlarmsNotificationManner{
			value: "TOPIC_SUBSCRIPTION",
		},
		NOTIFICATION_POLICY: ListAlarmsRespAlarmsNotificationManner{
			value: "NOTIFICATION_POLICY",
		},
	}
}

func (c ListAlarmsRespAlarmsNotificationManner) Value() string {
	return c.value
}

func (c ListAlarmsRespAlarmsNotificationManner) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmsRespAlarmsNotificationManner) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
