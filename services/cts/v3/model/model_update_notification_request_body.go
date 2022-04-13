package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 修改关键操作通知规则的请求体。
type UpdateNotificationRequestBody struct {
	// 标识关键操作名称。

	NotificationName string `json:"notification_name"`
	// 标识操作类型。 目前支持的操作类型有完整类型(complete)和自定义类型(customized)。 完整类型下，CTS发送通知的对象为已对接服务的所有事件。 自定义类型下，CTS发送通知的对象是在operations列表中指定的事件。

	OperationType UpdateNotificationRequestBodyOperationType `json:"operation_type"`
	// 操作事件列表。

	Operations *[]Operations `json:"operations,omitempty"`
	// 通知用户列表，目前最多支持对10个用户组和50个用户发起的操作进行配置。

	NotifyUserList *[]NotificationUsers `json:"notify_user_list,omitempty"`
	// 标识关键操作通知状态，包括正常(enabled)，停止(disabled)两种状态。

	Status UpdateNotificationRequestBodyStatus `json:"status"`
	// 消息通知服务的topic_urn或者函数工作流的func_urn，当“status”字段为enabled时，该字段必填。 - 消息通知服务的topic_urn可以通过消息通知服务的查询主题列表API获取，示例：urn:smn:regionId:f96188c7ccaf4ffba0c9aa149ab2bd57:test_topic_v2。 - 函数工作流的func_urn可以通过函数工作流的获取函数列表API获取，示例：urn:fss:xxxxxxxxx:7aad83af3e8d42e99ac194e8419e2c9b:function:default:test。

	TopicId *string `json:"topic_id,omitempty"`
	// 关键操作通知id

	NotificationId string `json:"notification_id"`
}

func (o UpdateNotificationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateNotificationRequestBody", string(data)}, " ")
}

type UpdateNotificationRequestBodyOperationType struct {
	value string
}

type UpdateNotificationRequestBodyOperationTypeEnum struct {
	CUSTOMIZED UpdateNotificationRequestBodyOperationType
	COMPLETE   UpdateNotificationRequestBodyOperationType
}

func GetUpdateNotificationRequestBodyOperationTypeEnum() UpdateNotificationRequestBodyOperationTypeEnum {
	return UpdateNotificationRequestBodyOperationTypeEnum{
		CUSTOMIZED: UpdateNotificationRequestBodyOperationType{
			value: "customized",
		},
		COMPLETE: UpdateNotificationRequestBodyOperationType{
			value: "complete",
		},
	}
}

func (c UpdateNotificationRequestBodyOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateNotificationRequestBodyOperationType) UnmarshalJSON(b []byte) error {
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

type UpdateNotificationRequestBodyStatus struct {
	value string
}

type UpdateNotificationRequestBodyStatusEnum struct {
	ENABLED  UpdateNotificationRequestBodyStatus
	DISABLED UpdateNotificationRequestBodyStatus
}

func GetUpdateNotificationRequestBodyStatusEnum() UpdateNotificationRequestBodyStatusEnum {
	return UpdateNotificationRequestBodyStatusEnum{
		ENABLED: UpdateNotificationRequestBodyStatus{
			value: "enabled",
		},
		DISABLED: UpdateNotificationRequestBodyStatus{
			value: "disabled",
		},
	}
}

func (c UpdateNotificationRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateNotificationRequestBodyStatus) UnmarshalJSON(b []byte) error {
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
