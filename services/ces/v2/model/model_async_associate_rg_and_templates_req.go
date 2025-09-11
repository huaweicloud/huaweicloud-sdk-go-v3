package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AsyncAssociateRgAndTemplatesReq struct {

	// 告警模板编号列表，当ID列表为空时，将删除该资源分组已关联的告警模板所创建的告警规则
	TemplateIds []string `json:"template_ids"`

	// **参数解释**： 是否开启告警通知。     **约束限制**： 不涉及。 **取值范围**： 布尔值。 - true:开启。 - false:关闭。 **默认取值**： true
	NotificationEnabled bool `json:"notification_enabled"`

	// 告警触发通知列表
	AlarmNotifications *[]Notification `json:"alarm_notifications,omitempty"`

	// 告警恢复通知列表
	OkNotifications *[]Notification `json:"ok_notifications,omitempty"`

	// **参数解释**： 告警通知开启时间。    **约束限制**： 不涉及。 **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。           **默认取值**： 不涉及。
	NotificationBeginTime *string `json:"notification_begin_time,omitempty"`

	// **参数解释**： 告警通知关闭时间。    **约束限制**： 不涉及。 **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。           **默认取值**： 不涉及。
	NotificationEndTime *string `json:"notification_end_time,omitempty"`

	// **参数解释**： 时区，形如：\"GMT-08:00\"、\"GMT+08:00\"、\"GMT+0:00\"。    **约束限制**： 不涉及。 **取值范围**： 长度为[1,16]个字符。           **默认取值**： 不涉及。
	EffectiveTimezone *string `json:"effective_timezone,omitempty"`

	// **参数解释**： 企业项目ID。     **约束限制**： 不涉及。 **取值范围**： 只能包含小写字母、数字、“-”。           **默认取值**： 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// NOTIFICATION_GROUP(通知组)/TOPIC_SUBSCRIPTION(主题订阅)/NOTIFICATION_POLICY(通知策略)
	NotificationManner *AsyncAssociateRgAndTemplatesReqNotificationManner `json:"notification_manner,omitempty"`

	// 关联的通知策略ID列表
	NotificationPolicyIds *[]string `json:"notification_policy_ids,omitempty"`
}

func (o AsyncAssociateRgAndTemplatesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AsyncAssociateRgAndTemplatesReq struct{}"
	}

	return strings.Join([]string{"AsyncAssociateRgAndTemplatesReq", string(data)}, " ")
}

type AsyncAssociateRgAndTemplatesReqNotificationManner struct {
	value string
}

type AsyncAssociateRgAndTemplatesReqNotificationMannerEnum struct {
	NOTIFICATION_GROUP  AsyncAssociateRgAndTemplatesReqNotificationManner
	TOPIC_SUBSCRIPTION  AsyncAssociateRgAndTemplatesReqNotificationManner
	NOTIFICATION_POLICY AsyncAssociateRgAndTemplatesReqNotificationManner
}

func GetAsyncAssociateRgAndTemplatesReqNotificationMannerEnum() AsyncAssociateRgAndTemplatesReqNotificationMannerEnum {
	return AsyncAssociateRgAndTemplatesReqNotificationMannerEnum{
		NOTIFICATION_GROUP: AsyncAssociateRgAndTemplatesReqNotificationManner{
			value: "NOTIFICATION_GROUP",
		},
		TOPIC_SUBSCRIPTION: AsyncAssociateRgAndTemplatesReqNotificationManner{
			value: "TOPIC_SUBSCRIPTION",
		},
		NOTIFICATION_POLICY: AsyncAssociateRgAndTemplatesReqNotificationManner{
			value: "NOTIFICATION_POLICY",
		},
	}
}

func (c AsyncAssociateRgAndTemplatesReqNotificationManner) Value() string {
	return c.value
}

func (c AsyncAssociateRgAndTemplatesReqNotificationManner) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AsyncAssociateRgAndTemplatesReqNotificationManner) UnmarshalJSON(b []byte) error {
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
