package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CreateProductResponse struct {

	// 权限
	Permissions *[]string `json:"permissions,omitempty" xml:"permissions"`

	// 产品ID
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 产品唯一序列（系统唯一值，用于MQS的TOPIC中标记产品）
	ProductSerial *string `json:"product_serial,omitempty" xml:"product_serial"`

	// 应用ID
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	// 产品名称，创建产品时租户内唯一，长度最大64，仅支持中文，英文字母，数字，下划线和中划线
	Name *string `json:"name,omitempty" xml:"name"`

	// 产品供应商ID
	ManufacturerId *string `json:"manufacturer_id,omitempty" xml:"manufacturer_id"`

	// 厂商名称
	ManufacturerName *string `json:"manufacturer_name,omitempty" xml:"manufacturer_name"`

	// 产品型号
	Model *string `json:"model,omitempty" xml:"model"`

	// 产品类型，0-普通产品(不支持子设备) 1-网关产品
	ProductType *CreateProductResponseProductType `json:"product_type,omitempty" xml:"product_type"`

	// 产品描述，长度0-200
	Description *string `json:"description,omitempty" xml:"description"`

	// 产品的协议类型 0-mqtt 2-modbus 4-opcua
	ProtocolType *CreateProductResponseProtocolType `json:"protocol_type,omitempty" xml:"protocol_type"`

	// 产品的设备类型（默认Default）
	DeviceType *string `json:"device_type,omitempty" xml:"device_type"`

	// 产品版本
	Version *string `json:"version,omitempty" xml:"version"`

	CreatedUser *CreatedUser `json:"created_user,omitempty" xml:"created_user"`

	LastUpdatedUser *LastUpdatedUser `json:"last_updated_user,omitempty" xml:"last_updated_user"`

	Authentication *Authentication `json:"authentication,omitempty" xml:"authentication"`

	// 创建时间，timestamp(ms)，使用UTC时区
	CreatedDatetime *int64 `json:"created_datetime,omitempty" xml:"created_datetime"`

	// 应用名称
	AppName *string `json:"app_name,omitempty" xml:"app_name"`

	// data_format 0-JSON 1-USER_DEFINED
	DataFormat     *int32 `json:"data_format,omitempty" xml:"data_format"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateProductResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProductResponse struct{}"
	}

	return strings.Join([]string{"CreateProductResponse", string(data)}, " ")
}

type CreateProductResponseProductType struct {
	value int32
}

type CreateProductResponseProductTypeEnum struct {
	E_0 CreateProductResponseProductType
	E_1 CreateProductResponseProductType
}

func GetCreateProductResponseProductTypeEnum() CreateProductResponseProductTypeEnum {
	return CreateProductResponseProductTypeEnum{
		E_0: CreateProductResponseProductType{
			value: 0,
		}, E_1: CreateProductResponseProductType{
			value: 1,
		},
	}
}

func (c CreateProductResponseProductType) Value() int32 {
	return c.value
}

func (c CreateProductResponseProductType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateProductResponseProductType) UnmarshalJSON(b []byte) error {
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

type CreateProductResponseProtocolType struct {
	value int32
}

type CreateProductResponseProtocolTypeEnum struct {
	E_0 CreateProductResponseProtocolType
	E_2 CreateProductResponseProtocolType
	E_4 CreateProductResponseProtocolType
}

func GetCreateProductResponseProtocolTypeEnum() CreateProductResponseProtocolTypeEnum {
	return CreateProductResponseProtocolTypeEnum{
		E_0: CreateProductResponseProtocolType{
			value: 0,
		}, E_2: CreateProductResponseProtocolType{
			value: 2,
		}, E_4: CreateProductResponseProtocolType{
			value: 4,
		},
	}
}

func (c CreateProductResponseProtocolType) Value() int32 {
	return c.value
}

func (c CreateProductResponseProtocolType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateProductResponseProtocolType) UnmarshalJSON(b []byte) error {
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
