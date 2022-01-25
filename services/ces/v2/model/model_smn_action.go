package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SmnAction struct {
	// 通知类型

	Type SmnActionType `json:"type"`
	// 告警状态发生变化时，被通知对象的列表。通知对象ID最多可以配置5个。topicUrn可从SMN获取，具体操作请参考查询Topic列表。当type为notification时，notificationList列表不能为空；当type为autoscaling时，列表必须为[]。 说明：若alarm_action_enabled为true，对应的alarm_actions、insufficientdata_actions（该参数已废弃，建议无需配置）、ok_actions至少有一个不能为空。若alarm_actions、insufficientdata_actions（该参数已废弃，建议无需配置）、ok_actions同时存在时，notificationList值保持一致。

	NotificationList []string `json:"notification_list"`
}

func (o SmnAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmnAction struct{}"
	}

	return strings.Join([]string{"SmnAction", string(data)}, " ")
}

type SmnActionType struct {
	value string
}

type SmnActionTypeEnum struct {
	NOTIFICATION  SmnActionType
	AUTOSCALING   SmnActionType
	GROUPWATCH    SmnActionType
	ECS_RECOVERY  SmnActionType
	CONTACT       SmnActionType
	CONTACT_GROUP SmnActionType
}

func GetSmnActionTypeEnum() SmnActionTypeEnum {
	return SmnActionTypeEnum{
		NOTIFICATION: SmnActionType{
			value: "notification",
		},
		AUTOSCALING: SmnActionType{
			value: "autoscaling",
		},
		GROUPWATCH: SmnActionType{
			value: "groupwatch",
		},
		ECS_RECOVERY: SmnActionType{
			value: "ecsRecovery",
		},
		CONTACT: SmnActionType{
			value: "contact",
		},
		CONTACT_GROUP: SmnActionType{
			value: "contactGroup",
		},
	}
}

func (c SmnActionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmnActionType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
