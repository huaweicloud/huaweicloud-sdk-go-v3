package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Response Object
type CreateSourceResponse struct {
	// 源数据源ID

	SourceId *int32 `json:"source_id,omitempty"`
	// 产品ID

	ProductId *int32 `json:"product_id,omitempty"`
	// 设备ID，不填为全部设备

	DeviceId *int32 `json:"device_id,omitempty"`
	// 主题，当设备ID为空时为产品级主题，设备ID不为空时为设备级主题

	Topic *string `json:"topic,omitempty"`
	// 设备名称

	DeviceName *string `json:"device_name,omitempty"`
	// 产品名称

	ProductName *string `json:"product_name,omitempty"`
	// 是否payload使用base64，0-是 1-否

	IsBase64 *CreateSourceResponseIsBase64 `json:"is_base64,omitempty"`
	// 是否包含设备信息，0-是，1-否

	ContainDeviceInfo *CreateSourceResponseContainDeviceInfo `json:"contain_device_info,omitempty"`
	HttpStatusCode    int                                    `json:"-"`
}

func (o CreateSourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSourceResponse struct{}"
	}

	return strings.Join([]string{"CreateSourceResponse", string(data)}, " ")
}

type CreateSourceResponseIsBase64 struct {
	value int32
}

type CreateSourceResponseIsBase64Enum struct {
	E_0 CreateSourceResponseIsBase64
	E_1 CreateSourceResponseIsBase64
}

func GetCreateSourceResponseIsBase64Enum() CreateSourceResponseIsBase64Enum {
	return CreateSourceResponseIsBase64Enum{
		E_0: CreateSourceResponseIsBase64{
			value: 0,
		}, E_1: CreateSourceResponseIsBase64{
			value: 1,
		},
	}
}

func (c CreateSourceResponseIsBase64) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSourceResponseIsBase64) UnmarshalJSON(b []byte) error {
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

type CreateSourceResponseContainDeviceInfo struct {
	value int32
}

type CreateSourceResponseContainDeviceInfoEnum struct {
	E_0 CreateSourceResponseContainDeviceInfo
	E_1 CreateSourceResponseContainDeviceInfo
}

func GetCreateSourceResponseContainDeviceInfoEnum() CreateSourceResponseContainDeviceInfoEnum {
	return CreateSourceResponseContainDeviceInfoEnum{
		E_0: CreateSourceResponseContainDeviceInfo{
			value: 0,
		}, E_1: CreateSourceResponseContainDeviceInfo{
			value: 1,
		},
	}
}

func (c CreateSourceResponseContainDeviceInfo) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSourceResponseContainDeviceInfo) UnmarshalJSON(b []byte) error {
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
