package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateNotificationRequestBody struct {

	// 通知归属的应用ID
	AppId string `json:"app_id" xml:"app_id"`

	// 通知类型 0-设备上线通知 1-设备下线通知 2-设备添加通知 3-设备删除通知 4-设备变更通知
	Type CreateNotificationRequestBodyType `json:"type" xml:"type"`

	// 通知发送的主题名，该主题需要在MQS存在
	Topic string `json:"topic" xml:"topic"`

	// 启停状态 0-启用 1-停用
	Status CreateNotificationRequestBodyStatus `json:"status" xml:"status"`
}

func (o CreateNotificationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotificationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateNotificationRequestBody", string(data)}, " ")
}

type CreateNotificationRequestBodyType struct {
	value int32
}

type CreateNotificationRequestBodyTypeEnum struct {
	E_0 CreateNotificationRequestBodyType
	E_1 CreateNotificationRequestBodyType
	E_2 CreateNotificationRequestBodyType
	E_3 CreateNotificationRequestBodyType
	E_4 CreateNotificationRequestBodyType
}

func GetCreateNotificationRequestBodyTypeEnum() CreateNotificationRequestBodyTypeEnum {
	return CreateNotificationRequestBodyTypeEnum{
		E_0: CreateNotificationRequestBodyType{
			value: 0,
		}, E_1: CreateNotificationRequestBodyType{
			value: 1,
		}, E_2: CreateNotificationRequestBodyType{
			value: 2,
		}, E_3: CreateNotificationRequestBodyType{
			value: 3,
		}, E_4: CreateNotificationRequestBodyType{
			value: 4,
		},
	}
}

func (c CreateNotificationRequestBodyType) Value() int32 {
	return c.value
}

func (c CreateNotificationRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNotificationRequestBodyType) UnmarshalJSON(b []byte) error {
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

type CreateNotificationRequestBodyStatus struct {
	value int32
}

type CreateNotificationRequestBodyStatusEnum struct {
	E_0 CreateNotificationRequestBodyStatus
	E_1 CreateNotificationRequestBodyStatus
}

func GetCreateNotificationRequestBodyStatusEnum() CreateNotificationRequestBodyStatusEnum {
	return CreateNotificationRequestBodyStatusEnum{
		E_0: CreateNotificationRequestBodyStatus{
			value: 0,
		}, E_1: CreateNotificationRequestBodyStatus{
			value: 1,
		},
	}
}

func (c CreateNotificationRequestBodyStatus) Value() int32 {
	return c.value
}

func (c CreateNotificationRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNotificationRequestBodyStatus) UnmarshalJSON(b []byte) error {
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
