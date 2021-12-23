package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CreateNotificationResponse struct {
	// 通知名称。

	NotificationName *string `json:"notification_name,omitempty"`
	// 操作类型，完整和自定义。

	OperationType *CreateNotificationResponseOperationType `json:"operation_type,omitempty"`
	// 操作事件列表。

	Operations *[]Operations `json:"operations,omitempty"`
	// 通知用户列表，目前最多支持对10个用户组和50个用户发起的操作进行配置。

	NotifyUserList *[]NotificationUsers `json:"notify_user_list,omitempty"`
	// 通知状态，启用和停用。

	Status *CreateNotificationResponseStatus `json:"status,omitempty"`
	// 消息通知服务(SMN)主题的唯一的资源标识，可通过查询主题列表获取该标识。

	TopicId *string `json:"topic_id,omitempty"`
	// 通知的唯一标识ID。

	NotificationId *string `json:"notification_id,omitempty"`
	// 通知类型，消息通知，函数触发器。

	NotificationType *CreateNotificationResponseNotificationType `json:"notification_type,omitempty"`
	// 项目ID。

	ProjectId *string `json:"project_id,omitempty"`
	// 通知规则创建时间。

	CreateTime     *int64 `json:"create_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateNotificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotificationResponse struct{}"
	}

	return strings.Join([]string{"CreateNotificationResponse", string(data)}, " ")
}

type CreateNotificationResponseOperationType struct {
	value string
}

type CreateNotificationResponseOperationTypeEnum struct {
	CUSTOMIZED CreateNotificationResponseOperationType
	COMPLETE   CreateNotificationResponseOperationType
}

func GetCreateNotificationResponseOperationTypeEnum() CreateNotificationResponseOperationTypeEnum {
	return CreateNotificationResponseOperationTypeEnum{
		CUSTOMIZED: CreateNotificationResponseOperationType{
			value: "customized",
		},
		COMPLETE: CreateNotificationResponseOperationType{
			value: "complete",
		},
	}
}

func (c CreateNotificationResponseOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNotificationResponseOperationType) UnmarshalJSON(b []byte) error {
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

type CreateNotificationResponseStatus struct {
	value string
}

type CreateNotificationResponseStatusEnum struct {
	ENABLED  CreateNotificationResponseStatus
	DISABLED CreateNotificationResponseStatus
}

func GetCreateNotificationResponseStatusEnum() CreateNotificationResponseStatusEnum {
	return CreateNotificationResponseStatusEnum{
		ENABLED: CreateNotificationResponseStatus{
			value: "enabled",
		},
		DISABLED: CreateNotificationResponseStatus{
			value: "disabled",
		},
	}
}

func (c CreateNotificationResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNotificationResponseStatus) UnmarshalJSON(b []byte) error {
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

type CreateNotificationResponseNotificationType struct {
	value string
}

type CreateNotificationResponseNotificationTypeEnum struct {
	SMN CreateNotificationResponseNotificationType
	FUN CreateNotificationResponseNotificationType
}

func GetCreateNotificationResponseNotificationTypeEnum() CreateNotificationResponseNotificationTypeEnum {
	return CreateNotificationResponseNotificationTypeEnum{
		SMN: CreateNotificationResponseNotificationType{
			value: "smn",
		},
		FUN: CreateNotificationResponseNotificationType{
			value: "fun",
		},
	}
}

func (c CreateNotificationResponseNotificationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNotificationResponseNotificationType) UnmarshalJSON(b []byte) error {
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
