package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateMessageNotificationPolicyResponse Response Object
type CreateMessageNotificationPolicyResponse struct {

	// id
	Id *string `json:"id,omitempty"`

	// 消息类型。job:任务执行结果。
	MessageType *CreateMessageNotificationPolicyResponseMessageType `json:"message_type,omitempty"`

	// 名称样式，用来匹配消息类型中所有符合该样式的消息。比如，message_type设置为job，name_pattern设置为ray_job*，表示匹配到所有以\"ray_job\"开头的job发出的消息。
	NamePattern *string `json:"name_pattern,omitempty"`

	// 通知类型。SUCCESS:成功通知；FAILED：失败通知
	NotificationTypes *[]CreateMessageNotificationPolicyResponseNotificationTypes `json:"notification_types,omitempty"`

	// 消息通知主题。
	TopicUrn       *string `json:"topic_urn,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMessageNotificationPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMessageNotificationPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateMessageNotificationPolicyResponse", string(data)}, " ")
}

type CreateMessageNotificationPolicyResponseMessageType struct {
	value string
}

type CreateMessageNotificationPolicyResponseMessageTypeEnum struct {
	JOB CreateMessageNotificationPolicyResponseMessageType
}

func GetCreateMessageNotificationPolicyResponseMessageTypeEnum() CreateMessageNotificationPolicyResponseMessageTypeEnum {
	return CreateMessageNotificationPolicyResponseMessageTypeEnum{
		JOB: CreateMessageNotificationPolicyResponseMessageType{
			value: "job",
		},
	}
}

func (c CreateMessageNotificationPolicyResponseMessageType) Value() string {
	return c.value
}

func (c CreateMessageNotificationPolicyResponseMessageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMessageNotificationPolicyResponseMessageType) UnmarshalJSON(b []byte) error {
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

type CreateMessageNotificationPolicyResponseNotificationTypes struct {
	value string
}

type CreateMessageNotificationPolicyResponseNotificationTypesEnum struct {
	SUCCESS CreateMessageNotificationPolicyResponseNotificationTypes
	FAILED  CreateMessageNotificationPolicyResponseNotificationTypes
}

func GetCreateMessageNotificationPolicyResponseNotificationTypesEnum() CreateMessageNotificationPolicyResponseNotificationTypesEnum {
	return CreateMessageNotificationPolicyResponseNotificationTypesEnum{
		SUCCESS: CreateMessageNotificationPolicyResponseNotificationTypes{
			value: "SUCCESS",
		},
		FAILED: CreateMessageNotificationPolicyResponseNotificationTypes{
			value: "FAILED",
		},
	}
}

func (c CreateMessageNotificationPolicyResponseNotificationTypes) Value() string {
	return c.value
}

func (c CreateMessageNotificationPolicyResponseNotificationTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMessageNotificationPolicyResponseNotificationTypes) UnmarshalJSON(b []byte) error {
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
