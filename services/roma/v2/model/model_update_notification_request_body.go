package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateNotificationRequestBody struct {
	// 通知发送的主题名，该主题需要在MQS存在

	Topic string `json:"topic"`
	// 启停状态 0-启用 1-停用

	Status UpdateNotificationRequestBodyStatus `json:"status"`
}

func (o UpdateNotificationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateNotificationRequestBody", string(data)}, " ")
}

type UpdateNotificationRequestBodyStatus struct {
	value int64
}

type UpdateNotificationRequestBodyStatusEnum struct {
	E_0 UpdateNotificationRequestBodyStatus
	E_1 UpdateNotificationRequestBodyStatus
}

func GetUpdateNotificationRequestBodyStatusEnum() UpdateNotificationRequestBodyStatusEnum {
	return UpdateNotificationRequestBodyStatusEnum{
		E_0: UpdateNotificationRequestBodyStatus{
			value: 0,
		}, E_1: UpdateNotificationRequestBodyStatus{
			value: 1,
		},
	}
}

func (c UpdateNotificationRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateNotificationRequestBodyStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int64")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int64)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int64 error")
	}
}
