package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Notification struct {

	// **参数解释**： 通知类型。 **约束限制**： 不涉及。 **取值范围**： 枚举值。 - notification：通知组或主题订阅。 - contact：云账号联系人。 - contactGroup：（已废弃）通知组。 - autoscaling：AS通知，只在AS中使用。 - groupwatch：已废弃，不推荐使用。 - ecsRecovery：已废弃，不推荐使用。 - iecAction：已废弃，不推荐使用。 **默认取值**： 不涉及。
	Type NotificationType `json:"type"`

	// **参数解释**：  告警状态发生变化时，被通知对象的列表。topicUrn可从SMN获取，具体操作请参考查询Topic列表。 **约束限制**： 当type为notification时，notification_list列表不能为空；当type为autoscaling时，列表必须为[]。若notification_enabled为true，对应的alarm_notifications、ok_notifications至少有一个不能为空。若alarm_notifications、ok_notification同时存在时，notification_list值保持一致。最多包含20个通知对象，最少可以为0个。
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
