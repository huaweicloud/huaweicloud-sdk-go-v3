package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateProductResponse struct {

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
	ProductType *UpdateProductResponseProductType `json:"product_type,omitempty" xml:"product_type"`

	// 产品描述，长度0-200
	Description *string `json:"description,omitempty" xml:"description"`

	// 产品的协议类型 0-mqtt 2-modbus 4-opcua
	ProtocolType *UpdateProductResponseProtocolType `json:"protocol_type,omitempty" xml:"protocol_type"`

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

func (o UpdateProductResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProductResponse struct{}"
	}

	return strings.Join([]string{"UpdateProductResponse", string(data)}, " ")
}

type UpdateProductResponseProductType struct {
	value int32
}

type UpdateProductResponseProductTypeEnum struct {
	E_0 UpdateProductResponseProductType
	E_1 UpdateProductResponseProductType
}

func GetUpdateProductResponseProductTypeEnum() UpdateProductResponseProductTypeEnum {
	return UpdateProductResponseProductTypeEnum{
		E_0: UpdateProductResponseProductType{
			value: 0,
		}, E_1: UpdateProductResponseProductType{
			value: 1,
		},
	}
}

func (c UpdateProductResponseProductType) Value() int32 {
	return c.value
}

func (c UpdateProductResponseProductType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateProductResponseProductType) UnmarshalJSON(b []byte) error {
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

type UpdateProductResponseProtocolType struct {
	value int32
}

type UpdateProductResponseProtocolTypeEnum struct {
	E_0 UpdateProductResponseProtocolType
	E_2 UpdateProductResponseProtocolType
	E_4 UpdateProductResponseProtocolType
}

func GetUpdateProductResponseProtocolTypeEnum() UpdateProductResponseProtocolTypeEnum {
	return UpdateProductResponseProtocolTypeEnum{
		E_0: UpdateProductResponseProtocolType{
			value: 0,
		}, E_2: UpdateProductResponseProtocolType{
			value: 2,
		}, E_4: UpdateProductResponseProtocolType{
			value: 4,
		},
	}
}

func (c UpdateProductResponseProtocolType) Value() int32 {
	return c.value
}

func (c UpdateProductResponseProtocolType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateProductResponseProtocolType) UnmarshalJSON(b []byte) error {
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
