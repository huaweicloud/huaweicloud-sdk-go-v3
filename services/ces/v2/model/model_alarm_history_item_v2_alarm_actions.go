package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AlarmHistoryItemV2AlarmActions struct {

	// **参数解释**： 通知类型。 **取值范围**： 枚举值。 - notification：通知组或主题订阅。 - contact：云账号联系人。 - contactGroup：（已废弃）通知组。 - autoscaling：AS通知，只在AS中使用。 - groupwatch：已废弃，不推荐使用。 - ecsRecovery：已废弃，不推荐使用。 - iecAction：已废弃，不推荐使用。
	Type *AlarmHistoryItemV2AlarmActionsType `json:"type,omitempty"`

	// **参数解释**： 告警状态发生变化时，被通知对象的列表。topicUrn可从SMN获取，具体操作请参考查询Topic列表。
	NotificationList *[]string `json:"notification_list,omitempty"`
}

func (o AlarmHistoryItemV2AlarmActions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmHistoryItemV2AlarmActions struct{}"
	}

	return strings.Join([]string{"AlarmHistoryItemV2AlarmActions", string(data)}, " ")
}

type AlarmHistoryItemV2AlarmActionsType struct {
	value string
}

type AlarmHistoryItemV2AlarmActionsTypeEnum struct {
	NOTIFICATION  AlarmHistoryItemV2AlarmActionsType
	AUTOSCALING   AlarmHistoryItemV2AlarmActionsType
	GROUPWATCH    AlarmHistoryItemV2AlarmActionsType
	ECS_RECOVERY  AlarmHistoryItemV2AlarmActionsType
	CONTACT       AlarmHistoryItemV2AlarmActionsType
	CONTACT_GROUP AlarmHistoryItemV2AlarmActionsType
	IEC_ACTION    AlarmHistoryItemV2AlarmActionsType
}

func GetAlarmHistoryItemV2AlarmActionsTypeEnum() AlarmHistoryItemV2AlarmActionsTypeEnum {
	return AlarmHistoryItemV2AlarmActionsTypeEnum{
		NOTIFICATION: AlarmHistoryItemV2AlarmActionsType{
			value: "notification",
		},
		AUTOSCALING: AlarmHistoryItemV2AlarmActionsType{
			value: "autoscaling",
		},
		GROUPWATCH: AlarmHistoryItemV2AlarmActionsType{
			value: "groupwatch",
		},
		ECS_RECOVERY: AlarmHistoryItemV2AlarmActionsType{
			value: "ecsRecovery",
		},
		CONTACT: AlarmHistoryItemV2AlarmActionsType{
			value: "contact",
		},
		CONTACT_GROUP: AlarmHistoryItemV2AlarmActionsType{
			value: "contactGroup",
		},
		IEC_ACTION: AlarmHistoryItemV2AlarmActionsType{
			value: "iecAction",
		},
	}
}

func (c AlarmHistoryItemV2AlarmActionsType) Value() string {
	return c.value
}

func (c AlarmHistoryItemV2AlarmActionsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmHistoryItemV2AlarmActionsType) UnmarshalJSON(b []byte) error {
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
