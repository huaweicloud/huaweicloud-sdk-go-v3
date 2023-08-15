package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Source struct {

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
	IsBase64 *SourceIsBase64 `json:"is_base64,omitempty"`

	// 是否包含设备信息，0-是，1-否
	ContainDeviceInfo *SourceContainDeviceInfo `json:"contain_device_info,omitempty"`
}

func (o Source) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Source struct{}"
	}

	return strings.Join([]string{"Source", string(data)}, " ")
}

type SourceIsBase64 struct {
	value int32
}

type SourceIsBase64Enum struct {
	E_0 SourceIsBase64
	E_1 SourceIsBase64
}

func GetSourceIsBase64Enum() SourceIsBase64Enum {
	return SourceIsBase64Enum{
		E_0: SourceIsBase64{
			value: 0,
		}, E_1: SourceIsBase64{
			value: 1,
		},
	}
}

func (c SourceIsBase64) Value() int32 {
	return c.value
}

func (c SourceIsBase64) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceIsBase64) UnmarshalJSON(b []byte) error {
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

type SourceContainDeviceInfo struct {
	value int32
}

type SourceContainDeviceInfoEnum struct {
	E_0 SourceContainDeviceInfo
	E_1 SourceContainDeviceInfo
}

func GetSourceContainDeviceInfoEnum() SourceContainDeviceInfoEnum {
	return SourceContainDeviceInfoEnum{
		E_0: SourceContainDeviceInfo{
			value: 0,
		}, E_1: SourceContainDeviceInfo{
			value: 1,
		},
	}
}

func (c SourceContainDeviceInfo) Value() int32 {
	return c.value
}

func (c SourceContainDeviceInfo) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceContainDeviceInfo) UnmarshalJSON(b []byte) error {
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
