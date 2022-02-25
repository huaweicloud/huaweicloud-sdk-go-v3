package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NotificationResponseBody struct {
	// 订阅ID

	NotificationId *int64 `json:"notification_id,omitempty"`
	// 订阅类型, 0:设备上线通知类型, 1:设备下线通知类型, 2:设备添加通知类型, 3:设备删除通知类型, 4:设备变更通知类型

	Type *NotificationResponseBodyType `json:"type,omitempty"`
	// 订阅管理状态，0：启用，1：停用

	Status *NotificationResponseBodyStatus `json:"status,omitempty"`
	// 订阅的topic名称

	Topic *string `json:"topic,omitempty"`
	// 实例ID

	InstanceId *string `json:"instance_id,omitempty"`
	// 应用ID

	AppId *string `json:"app_id,omitempty"`
}

func (o NotificationResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationResponseBody struct{}"
	}

	return strings.Join([]string{"NotificationResponseBody", string(data)}, " ")
}

type NotificationResponseBodyType struct {
	value int32
}

type NotificationResponseBodyTypeEnum struct {
	E_0 NotificationResponseBodyType
	E_1 NotificationResponseBodyType
	E_2 NotificationResponseBodyType
	E_3 NotificationResponseBodyType
	E_4 NotificationResponseBodyType
}

func GetNotificationResponseBodyTypeEnum() NotificationResponseBodyTypeEnum {
	return NotificationResponseBodyTypeEnum{
		E_0: NotificationResponseBodyType{
			value: 0,
		}, E_1: NotificationResponseBodyType{
			value: 1,
		}, E_2: NotificationResponseBodyType{
			value: 2,
		}, E_3: NotificationResponseBodyType{
			value: 3,
		}, E_4: NotificationResponseBodyType{
			value: 4,
		},
	}
}

func (c NotificationResponseBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotificationResponseBodyType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type NotificationResponseBodyStatus struct {
	value int32
}

type NotificationResponseBodyStatusEnum struct {
	E_0 NotificationResponseBodyStatus
	E_1 NotificationResponseBodyStatus
}

func GetNotificationResponseBodyStatusEnum() NotificationResponseBodyStatusEnum {
	return NotificationResponseBodyStatusEnum{
		E_0: NotificationResponseBodyStatus{
			value: 0,
		}, E_1: NotificationResponseBodyStatus{
			value: 1,
		},
	}
}

func (c NotificationResponseBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotificationResponseBodyStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
