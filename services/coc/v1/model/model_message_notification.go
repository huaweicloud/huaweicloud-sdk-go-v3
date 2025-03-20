package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MessageNotification 消息通知
type MessageNotification struct {

	// 通知策略
	Policy *string `json:"policy,omitempty"`

	// 通知对象类型
	NotificationEndpointType MessageNotificationNotificationEndpointType `json:"notification_endpoint_type"`

	// 排班场景ID
	ScheduleSceneId *string `json:"schedule_scene_id,omitempty"`

	// 排班角色ID
	ScheduleRoleId *string `json:"schedule_role_id,omitempty"`

	// 消息接收人
	Recipients *string `json:"recipients,omitempty"`

	// 通知渠道
	Protocol *MessageNotificationProtocol `json:"protocol,omitempty"`
}

func (o MessageNotification) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MessageNotification struct{}"
	}

	return strings.Join([]string{"MessageNotification", string(data)}, " ")
}

type MessageNotificationNotificationEndpointType struct {
	value string
}

type MessageNotificationNotificationEndpointTypeEnum struct {
	USER    MessageNotificationNotificationEndpointType
	ON_CALL MessageNotificationNotificationEndpointType
}

func GetMessageNotificationNotificationEndpointTypeEnum() MessageNotificationNotificationEndpointTypeEnum {
	return MessageNotificationNotificationEndpointTypeEnum{
		USER: MessageNotificationNotificationEndpointType{
			value: "USER",
		},
		ON_CALL: MessageNotificationNotificationEndpointType{
			value: "ON_CALL",
		},
	}
}

func (c MessageNotificationNotificationEndpointType) Value() string {
	return c.value
}

func (c MessageNotificationNotificationEndpointType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MessageNotificationNotificationEndpointType) UnmarshalJSON(b []byte) error {
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

type MessageNotificationProtocol struct {
	value string
}

type MessageNotificationProtocolEnum struct {
	DEFAULT    MessageNotificationProtocol
	NONE       MessageNotificationProtocol
	SMS        MessageNotificationProtocol
	EMAIL      MessageNotificationProtocol
	DINGDING   MessageNotificationProtocol
	LARK       MessageNotificationProtocol
	WELINK     MessageNotificationProtocol
	CALLNOTIFY MessageNotificationProtocol
	WECHAT     MessageNotificationProtocol
}

func GetMessageNotificationProtocolEnum() MessageNotificationProtocolEnum {
	return MessageNotificationProtocolEnum{
		DEFAULT: MessageNotificationProtocol{
			value: "DEFAULT",
		},
		NONE: MessageNotificationProtocol{
			value: "NONE",
		},
		SMS: MessageNotificationProtocol{
			value: "SMS",
		},
		EMAIL: MessageNotificationProtocol{
			value: "EMAIL",
		},
		DINGDING: MessageNotificationProtocol{
			value: "DINGDING",
		},
		LARK: MessageNotificationProtocol{
			value: "LARK",
		},
		WELINK: MessageNotificationProtocol{
			value: "WELINK",
		},
		CALLNOTIFY: MessageNotificationProtocol{
			value: "CALLNOTIFY",
		},
		WECHAT: MessageNotificationProtocol{
			value: "WECHAT",
		},
	}
}

func (c MessageNotificationProtocol) Value() string {
	return c.value
}

func (c MessageNotificationProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MessageNotificationProtocol) UnmarshalJSON(b []byte) error {
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
