package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateNotificationResponse struct {

	// 订阅ID
	NotificationId *int64 `json:"notification_id,omitempty"`

	// 订阅类型, 0:设备上线通知类型, 1:设备下线通知类型, 2:设备添加通知类型, 3:设备删除通知类型, 4:设备变更通知类型
	Type *UpdateNotificationResponseType `json:"type,omitempty"`

	// 订阅管理状态，0：启用，1：停用
	Status *UpdateNotificationResponseStatus `json:"status,omitempty"`

	// 订阅的topic名称
	Topic *string `json:"topic,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 应用ID
	AppId          *string `json:"app_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateNotificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationResponse struct{}"
	}

	return strings.Join([]string{"UpdateNotificationResponse", string(data)}, " ")
}

type UpdateNotificationResponseType struct {
	value int32
}

type UpdateNotificationResponseTypeEnum struct {
	E_0 UpdateNotificationResponseType
	E_1 UpdateNotificationResponseType
	E_2 UpdateNotificationResponseType
	E_3 UpdateNotificationResponseType
	E_4 UpdateNotificationResponseType
}

func GetUpdateNotificationResponseTypeEnum() UpdateNotificationResponseTypeEnum {
	return UpdateNotificationResponseTypeEnum{
		E_0: UpdateNotificationResponseType{
			value: 0,
		}, E_1: UpdateNotificationResponseType{
			value: 1,
		}, E_2: UpdateNotificationResponseType{
			value: 2,
		}, E_3: UpdateNotificationResponseType{
			value: 3,
		}, E_4: UpdateNotificationResponseType{
			value: 4,
		},
	}
}

func (c UpdateNotificationResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateNotificationResponseType) UnmarshalJSON(b []byte) error {
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

type UpdateNotificationResponseStatus struct {
	value int32
}

type UpdateNotificationResponseStatusEnum struct {
	E_0 UpdateNotificationResponseStatus
	E_1 UpdateNotificationResponseStatus
}

func GetUpdateNotificationResponseStatusEnum() UpdateNotificationResponseStatusEnum {
	return UpdateNotificationResponseStatusEnum{
		E_0: UpdateNotificationResponseStatus{
			value: 0,
		}, E_1: UpdateNotificationResponseStatus{
			value: 1,
		},
	}
}

func (c UpdateNotificationResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateNotificationResponseStatus) UnmarshalJSON(b []byte) error {
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
