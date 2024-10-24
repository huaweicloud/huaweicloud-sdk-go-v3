package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateMessageNotificationPolicy struct {

	// 消息类型。job:任务执行结果。
	MessageType CreateMessageNotificationPolicyMessageType `json:"message_type"`

	// 名称样式，用来匹配消息类型中所有符合该样式的消息。比如，message_type设置为job，name_pattern设置为ray_job*，表示匹配到所有以\"ray_job\"开头的job发出的消息。
	NamePattern string `json:"name_pattern"`

	// 通知类型。SUCCESS:成功通知；FAILED：失败通知
	NotificationTypes []CreateMessageNotificationPolicyNotificationTypes `json:"notification_types"`

	// 消息通知主题。
	TopicUrn string `json:"topic_urn"`
}

func (o CreateMessageNotificationPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMessageNotificationPolicy struct{}"
	}

	return strings.Join([]string{"CreateMessageNotificationPolicy", string(data)}, " ")
}

type CreateMessageNotificationPolicyMessageType struct {
	value string
}

type CreateMessageNotificationPolicyMessageTypeEnum struct {
	JOB CreateMessageNotificationPolicyMessageType
}

func GetCreateMessageNotificationPolicyMessageTypeEnum() CreateMessageNotificationPolicyMessageTypeEnum {
	return CreateMessageNotificationPolicyMessageTypeEnum{
		JOB: CreateMessageNotificationPolicyMessageType{
			value: "job",
		},
	}
}

func (c CreateMessageNotificationPolicyMessageType) Value() string {
	return c.value
}

func (c CreateMessageNotificationPolicyMessageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMessageNotificationPolicyMessageType) UnmarshalJSON(b []byte) error {
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

type CreateMessageNotificationPolicyNotificationTypes struct {
	value string
}

type CreateMessageNotificationPolicyNotificationTypesEnum struct {
	SUCCESS CreateMessageNotificationPolicyNotificationTypes
	FAILED  CreateMessageNotificationPolicyNotificationTypes
}

func GetCreateMessageNotificationPolicyNotificationTypesEnum() CreateMessageNotificationPolicyNotificationTypesEnum {
	return CreateMessageNotificationPolicyNotificationTypesEnum{
		SUCCESS: CreateMessageNotificationPolicyNotificationTypes{
			value: "SUCCESS",
		},
		FAILED: CreateMessageNotificationPolicyNotificationTypes{
			value: "FAILED",
		},
	}
}

func (c CreateMessageNotificationPolicyNotificationTypes) Value() string {
	return c.value
}

func (c CreateMessageNotificationPolicyNotificationTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMessageNotificationPolicyNotificationTypes) UnmarshalJSON(b []byte) error {
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
