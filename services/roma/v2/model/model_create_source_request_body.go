package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateSourceRequestBody struct {

	// 产品ID，自动向下取整
	ProductId int32 `json:"product_id"`

	// 设备ID，自动向下取整，不填为全部设备
	DeviceId *int32 `json:"device_id,omitempty"`

	// 主题，当设备ID为空时为产品级主题，设备ID不为空时为设备级主题
	Topic string `json:"topic"`

	// 是否payload使用base64，0-是 1-否
	IsBase64 *CreateSourceRequestBodyIsBase64 `json:"is_base64,omitempty"`

	// 是否包含设备信息是否包含设备信息，0-是 1-否
	ContainDeviceInfo *CreateSourceRequestBodyContainDeviceInfo `json:"contain_device_info,omitempty"`
}

func (o CreateSourceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSourceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSourceRequestBody", string(data)}, " ")
}

type CreateSourceRequestBodyIsBase64 struct {
	value int32
}

type CreateSourceRequestBodyIsBase64Enum struct {
	E_0 CreateSourceRequestBodyIsBase64
	E_1 CreateSourceRequestBodyIsBase64
}

func GetCreateSourceRequestBodyIsBase64Enum() CreateSourceRequestBodyIsBase64Enum {
	return CreateSourceRequestBodyIsBase64Enum{
		E_0: CreateSourceRequestBodyIsBase64{
			value: 0,
		}, E_1: CreateSourceRequestBodyIsBase64{
			value: 1,
		},
	}
}

func (c CreateSourceRequestBodyIsBase64) Value() int32 {
	return c.value
}

func (c CreateSourceRequestBodyIsBase64) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSourceRequestBodyIsBase64) UnmarshalJSON(b []byte) error {
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

type CreateSourceRequestBodyContainDeviceInfo struct {
	value int32
}

type CreateSourceRequestBodyContainDeviceInfoEnum struct {
	E_0 CreateSourceRequestBodyContainDeviceInfo
	E_1 CreateSourceRequestBodyContainDeviceInfo
}

func GetCreateSourceRequestBodyContainDeviceInfoEnum() CreateSourceRequestBodyContainDeviceInfoEnum {
	return CreateSourceRequestBodyContainDeviceInfoEnum{
		E_0: CreateSourceRequestBodyContainDeviceInfo{
			value: 0,
		}, E_1: CreateSourceRequestBodyContainDeviceInfo{
			value: 1,
		},
	}
}

func (c CreateSourceRequestBodyContainDeviceInfo) Value() int32 {
	return c.value
}

func (c CreateSourceRequestBodyContainDeviceInfo) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSourceRequestBodyContainDeviceInfo) UnmarshalJSON(b []byte) error {
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
