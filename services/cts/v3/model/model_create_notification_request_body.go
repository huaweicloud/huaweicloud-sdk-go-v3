package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建通知规则的请求体
type CreateNotificationRequestBody struct {
	// 标识关键操作名称。

	NotificationName string `json:"notification_name"`
	// 标识操作类型。 目前支持的操作类型有完整类型(complete)和自定义类型(customized)。 完整类型下，CTS发送通知的对象为已对接服务的所有事件，此时不用指定operations和notify_user_list字段。 自定义类型下，CTS发送通知的对象是在operations列表中指定的事件。

	OperationType CreateNotificationRequestBodyOperationType `json:"operation_type"`
	// 操作事件列表。

	Operations *[]Operations `json:"operations,omitempty"`
	// 通知用户列表，目前最多支持对10个用户组和50个用户发起的操作进行配置。

	NotifyUserList *[]NotificationUsers `json:"notify_user_list,omitempty"`
	// 消息通知服务的topic_urn或者函数工作流的func_urn。 - 消息通知服务的topic_urn可以通过消息通知服务的查询主题列表API获取，示例：urn:smn:regionId:f96188c7ccaf4ffba0c9aa149ab2bd57:test_topic_v2。 - 函数工作流的func_urn可以通过函数工作流的获取函数列表API获取，示例：urn:fss:xxxxxxxxx:7aad83af3e8d42e99ac194e8419e2c9b:function:default:test。

	TopicId *string `json:"topic_id,omitempty"`
}

func (o CreateNotificationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotificationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateNotificationRequestBody", string(data)}, " ")
}

type CreateNotificationRequestBodyOperationType struct {
	value string
}

type CreateNotificationRequestBodyOperationTypeEnum struct {
	COMPLETE   CreateNotificationRequestBodyOperationType
	CUSTOMIZED CreateNotificationRequestBodyOperationType
}

func GetCreateNotificationRequestBodyOperationTypeEnum() CreateNotificationRequestBodyOperationTypeEnum {
	return CreateNotificationRequestBodyOperationTypeEnum{
		COMPLETE: CreateNotificationRequestBodyOperationType{
			value: "complete",
		},
		CUSTOMIZED: CreateNotificationRequestBodyOperationType{
			value: "customized",
		},
	}
}

func (c CreateNotificationRequestBodyOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNotificationRequestBodyOperationType) UnmarshalJSON(b []byte) error {
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
