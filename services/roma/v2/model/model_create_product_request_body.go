package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateProductRequestBody struct {

	// 应用ID
	AppId string `json:"app_id"`

	// 产品名称，创建产品时租户内唯一，长度最大64，仅支持中文，英文字母，数字，下划线和中划线
	Name string `json:"name"`

	// 产品供应商ID
	ManufacturerId string `json:"manufacturer_id"`

	// 厂商名称
	ManufacturerName string `json:"manufacturer_name"`

	// 产品型号
	Model string `json:"model"`

	// 产品类型，0-普通产品(不支持子设备) 1-网关产品
	ProductType CreateProductRequestBodyProductType `json:"product_type"`

	// 产品描述，长度0-200
	Description *string `json:"description,omitempty"`

	// 产品的协议类型 0-mqtt 2-modbus 4-opcua
	ProtocolType CreateProductRequestBodyProtocolType `json:"protocol_type"`

	// 产品的设备类型（默认Default）
	DeviceType *string `json:"device_type,omitempty"`

	// 关联产品模板ID（使用产品模板创建产品时使用，否则为空），自动向下取整
	TemplateId *int32 `json:"template_id,omitempty"`

	// 模型版本
	Version *string `json:"version,omitempty"`

	// 产品的数据格式 0-JSON 1-USER_DEFINED
	DataFormat *CreateProductRequestBodyDataFormat `json:"data_format,omitempty"`
}

func (o CreateProductRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProductRequestBody struct{}"
	}

	return strings.Join([]string{"CreateProductRequestBody", string(data)}, " ")
}

type CreateProductRequestBodyProductType struct {
	value int32
}

type CreateProductRequestBodyProductTypeEnum struct {
	E_0 CreateProductRequestBodyProductType
	E_1 CreateProductRequestBodyProductType
}

func GetCreateProductRequestBodyProductTypeEnum() CreateProductRequestBodyProductTypeEnum {
	return CreateProductRequestBodyProductTypeEnum{
		E_0: CreateProductRequestBodyProductType{
			value: 0,
		}, E_1: CreateProductRequestBodyProductType{
			value: 1,
		},
	}
}

func (c CreateProductRequestBodyProductType) Value() int32 {
	return c.value
}

func (c CreateProductRequestBodyProductType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateProductRequestBodyProductType) UnmarshalJSON(b []byte) error {
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

type CreateProductRequestBodyProtocolType struct {
	value int32
}

type CreateProductRequestBodyProtocolTypeEnum struct {
	E_0 CreateProductRequestBodyProtocolType
	E_2 CreateProductRequestBodyProtocolType
	E_4 CreateProductRequestBodyProtocolType
}

func GetCreateProductRequestBodyProtocolTypeEnum() CreateProductRequestBodyProtocolTypeEnum {
	return CreateProductRequestBodyProtocolTypeEnum{
		E_0: CreateProductRequestBodyProtocolType{
			value: 0,
		}, E_2: CreateProductRequestBodyProtocolType{
			value: 2,
		}, E_4: CreateProductRequestBodyProtocolType{
			value: 4,
		},
	}
}

func (c CreateProductRequestBodyProtocolType) Value() int32 {
	return c.value
}

func (c CreateProductRequestBodyProtocolType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateProductRequestBodyProtocolType) UnmarshalJSON(b []byte) error {
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

type CreateProductRequestBodyDataFormat struct {
	value int32
}

type CreateProductRequestBodyDataFormatEnum struct {
	E_0 CreateProductRequestBodyDataFormat
	E_1 CreateProductRequestBodyDataFormat
}

func GetCreateProductRequestBodyDataFormatEnum() CreateProductRequestBodyDataFormatEnum {
	return CreateProductRequestBodyDataFormatEnum{
		E_0: CreateProductRequestBodyDataFormat{
			value: 0,
		}, E_1: CreateProductRequestBodyDataFormat{
			value: 1,
		},
	}
}

func (c CreateProductRequestBodyDataFormat) Value() int32 {
	return c.value
}

func (c CreateProductRequestBodyDataFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateProductRequestBodyDataFormat) UnmarshalJSON(b []byte) error {
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
