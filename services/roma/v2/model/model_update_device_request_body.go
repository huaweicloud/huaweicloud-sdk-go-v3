package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateDeviceRequestBody struct {

	// 设备名称，支持中文、中文标点符号（）。；，：“”、？《》及英文大小写、数字及英文符号()_,#.?'-@%&!, 长度2-64
	DeviceName string `json:"device_name"`

	// 设备状态 0启用 1禁用
	Status UpdateDeviceRequestBodyStatus `json:"status"`

	// 备注
	Description *string `json:"description,omitempty"`

	// 标签
	Tags *[]string `json:"tags,omitempty"`
}

func (o UpdateDeviceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDeviceRequestBody", string(data)}, " ")
}

type UpdateDeviceRequestBodyStatus struct {
	value int32
}

type UpdateDeviceRequestBodyStatusEnum struct {
	E_0 UpdateDeviceRequestBodyStatus
	E_1 UpdateDeviceRequestBodyStatus
}

func GetUpdateDeviceRequestBodyStatusEnum() UpdateDeviceRequestBodyStatusEnum {
	return UpdateDeviceRequestBodyStatusEnum{
		E_0: UpdateDeviceRequestBodyStatus{
			value: 0,
		}, E_1: UpdateDeviceRequestBodyStatus{
			value: 1,
		},
	}
}

func (c UpdateDeviceRequestBodyStatus) Value() int32 {
	return c.value
}

func (c UpdateDeviceRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDeviceRequestBodyStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
