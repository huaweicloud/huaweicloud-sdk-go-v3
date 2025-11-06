package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NotificationResp struct {

	// **参数解释**： 通知类型。 **取值范围**： 枚举值。 - notification：通知组或主题订阅。 - contact：云账号联系人。 - contactGroup：（已废弃）通知组。 - autoscaling：AS通知，只在AS中使用，不推荐客户使用。 - groupwatch：已废弃，不推荐使用。 - ecsRecovery：已废弃，不推荐使用。 - iecAction：已废弃，不推荐使用。
	Type *NotificationRespType `json:"type,omitempty"`

	// **参数解释**：  告警状态发生变化时，被通知对象的列表。topicUrn可从SMN获取，具体操作请参考查询Topic列表。
	NotificationList *[]string `json:"notification_list,omitempty"`
}

func (o NotificationResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationResp struct{}"
	}

	return strings.Join([]string{"NotificationResp", string(data)}, " ")
}

type NotificationRespType struct {
	value string
}

type NotificationRespTypeEnum struct {
	NOTIFICATION  NotificationRespType
	AUTOSCALING   NotificationRespType
	GROUPWATCH    NotificationRespType
	ECS_RECOVERY  NotificationRespType
	CONTACT       NotificationRespType
	CONTACT_GROUP NotificationRespType
	IEC_ACTION    NotificationRespType
}

func GetNotificationRespTypeEnum() NotificationRespTypeEnum {
	return NotificationRespTypeEnum{
		NOTIFICATION: NotificationRespType{
			value: "notification",
		},
		AUTOSCALING: NotificationRespType{
			value: "autoscaling",
		},
		GROUPWATCH: NotificationRespType{
			value: "groupwatch",
		},
		ECS_RECOVERY: NotificationRespType{
			value: "ecsRecovery",
		},
		CONTACT: NotificationRespType{
			value: "contact",
		},
		CONTACT_GROUP: NotificationRespType{
			value: "contactGroup",
		},
		IEC_ACTION: NotificationRespType{
			value: "iecAction",
		},
	}
}

func (c NotificationRespType) Value() string {
	return c.value
}

func (c NotificationRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotificationRespType) UnmarshalJSON(b []byte) error {
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
