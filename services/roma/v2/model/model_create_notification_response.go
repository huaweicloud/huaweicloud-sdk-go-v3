package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CreateNotificationResponse struct {

	// 订阅ID
	NotificationId *int64 `json:"notification_id,omitempty" xml:"notification_id"`

	// 订阅类型, 0:设备上线通知类型, 1:设备下线通知类型, 2:设备添加通知类型, 3:设备删除通知类型, 4:设备变更通知类型
	Type *CreateNotificationResponseType `json:"type,omitempty" xml:"type"`

	// 订阅管理状态，0：启用，1：停用
	Status *CreateNotificationResponseStatus `json:"status,omitempty" xml:"status"`

	// 订阅的topic名称
	Topic *string `json:"topic,omitempty" xml:"topic"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 应用ID
	AppId          *string `json:"app_id,omitempty" xml:"app_id"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateNotificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotificationResponse struct{}"
	}

	return strings.Join([]string{"CreateNotificationResponse", string(data)}, " ")
}

type CreateNotificationResponseType struct {
	value int32
}

type CreateNotificationResponseTypeEnum struct {
	E_0 CreateNotificationResponseType
	E_1 CreateNotificationResponseType
	E_2 CreateNotificationResponseType
	E_3 CreateNotificationResponseType
	E_4 CreateNotificationResponseType
}

func GetCreateNotificationResponseTypeEnum() CreateNotificationResponseTypeEnum {
	return CreateNotificationResponseTypeEnum{
		E_0: CreateNotificationResponseType{
			value: 0,
		}, E_1: CreateNotificationResponseType{
			value: 1,
		}, E_2: CreateNotificationResponseType{
			value: 2,
		}, E_3: CreateNotificationResponseType{
			value: 3,
		}, E_4: CreateNotificationResponseType{
			value: 4,
		},
	}
}

func (c CreateNotificationResponseType) Value() int32 {
	return c.value
}

func (c CreateNotificationResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNotificationResponseType) UnmarshalJSON(b []byte) error {
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

type CreateNotificationResponseStatus struct {
	value int32
}

type CreateNotificationResponseStatusEnum struct {
	E_0 CreateNotificationResponseStatus
	E_1 CreateNotificationResponseStatus
}

func GetCreateNotificationResponseStatusEnum() CreateNotificationResponseStatusEnum {
	return CreateNotificationResponseStatusEnum{
		E_0: CreateNotificationResponseStatus{
			value: 0,
		}, E_1: CreateNotificationResponseStatus{
			value: 1,
		},
	}
}

func (c CreateNotificationResponseStatus) Value() int32 {
	return c.value
}

func (c CreateNotificationResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNotificationResponseStatus) UnmarshalJSON(b []byte) error {
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
