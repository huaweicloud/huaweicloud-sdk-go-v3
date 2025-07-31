package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Notification struct {

	// 通知类型。notification为SMN通知，contact为云账号联系人，contactGroup为通知组。autoscaling为AS通知，只在AS中使用，不推荐客户使用。groupwatch、ecsRecovery及iecAction，已废弃，不推荐使用。
	Type NotificationType `json:"type"`

	// 告警状态发生变化时，被通知对象的列表。topicUrn可从SMN获取，具体操作请参考查询Topic列表。当type为notification时，notification_list列表不能为空；当type为autoscaling时，列表必须为[]。 说明：若alarm_action_enabled为true，对应的alarm_actions、ok_actions至少有一个不能为空。若alarm_actions、ok_actions同时存在时，notification_list值保持一致。
	NotificationList []string `json:"notification_list"`
}

func (o Notification) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Notification struct{}"
	}

	return strings.Join([]string{"Notification", string(data)}, " ")
}

type NotificationType struct {
	value string
}

type NotificationTypeEnum struct {
	NOTIFICATION  NotificationType
	AUTOSCALING   NotificationType
	GROUPWATCH    NotificationType
	ECS_RECOVERY  NotificationType
	CONTACT       NotificationType
	CONTACT_GROUP NotificationType
	IEC_ACTION    NotificationType
}

func GetNotificationTypeEnum() NotificationTypeEnum {
	return NotificationTypeEnum{
		NOTIFICATION: NotificationType{
			value: "notification",
		},
		AUTOSCALING: NotificationType{
			value: "autoscaling",
		},
		GROUPWATCH: NotificationType{
			value: "groupwatch",
		},
		ECS_RECOVERY: NotificationType{
			value: "ecsRecovery",
		},
		CONTACT: NotificationType{
			value: "contact",
		},
		CONTACT_GROUP: NotificationType{
			value: "contactGroup",
		},
		IEC_ACTION: NotificationType{
			value: "iecAction",
		},
	}
}

func (c NotificationType) Value() string {
	return c.value
}

func (c NotificationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotificationType) UnmarshalJSON(b []byte) error {
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
