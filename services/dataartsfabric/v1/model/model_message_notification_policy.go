package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MessageNotificationPolicy 消息通知策略
type MessageNotificationPolicy struct {

	// id
	Id string `json:"id"`

	// 消息类型。job:任务执行结果。
	MessageType MessageNotificationPolicyMessageType `json:"message_type"`

	// 名称样式，用来匹配消息类型中所有符合该样式的消息。比如，message_type设置为job，name_pattern设置为ray_job*，表示匹配到所有以\"ray_job\"开头的job发出的消息。
	NamePattern string `json:"name_pattern"`

	// 通知类型。SUCCESS:成功通知；FAILED：失败通知
	NotificationTypes []MessageNotificationPolicyNotificationTypes `json:"notification_types"`

	// 消息通知主题。
	TopicUrn string `json:"topic_urn"`
}

func (o MessageNotificationPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MessageNotificationPolicy struct{}"
	}

	return strings.Join([]string{"MessageNotificationPolicy", string(data)}, " ")
}

type MessageNotificationPolicyMessageType struct {
	value string
}

type MessageNotificationPolicyMessageTypeEnum struct {
	JOB MessageNotificationPolicyMessageType
}

func GetMessageNotificationPolicyMessageTypeEnum() MessageNotificationPolicyMessageTypeEnum {
	return MessageNotificationPolicyMessageTypeEnum{
		JOB: MessageNotificationPolicyMessageType{
			value: "job",
		},
	}
}

func (c MessageNotificationPolicyMessageType) Value() string {
	return c.value
}

func (c MessageNotificationPolicyMessageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MessageNotificationPolicyMessageType) UnmarshalJSON(b []byte) error {
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

type MessageNotificationPolicyNotificationTypes struct {
	value string
}

type MessageNotificationPolicyNotificationTypesEnum struct {
	SUCCESS MessageNotificationPolicyNotificationTypes
	FAILED  MessageNotificationPolicyNotificationTypes
}

func GetMessageNotificationPolicyNotificationTypesEnum() MessageNotificationPolicyNotificationTypesEnum {
	return MessageNotificationPolicyNotificationTypesEnum{
		SUCCESS: MessageNotificationPolicyNotificationTypes{
			value: "SUCCESS",
		},
		FAILED: MessageNotificationPolicyNotificationTypes{
			value: "FAILED",
		},
	}
}

func (c MessageNotificationPolicyNotificationTypes) Value() string {
	return c.value
}

func (c MessageNotificationPolicyNotificationTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MessageNotificationPolicyNotificationTypes) UnmarshalJSON(b []byte) error {
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
