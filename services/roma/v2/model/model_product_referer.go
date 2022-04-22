package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 产品
type ProductReferer struct {

	// 产品ID，未填写厂商ID+型号时产品ID必填
	ProductId *int32 `json:"product_id,omitempty"`

	// 产品名称
	ProductName *string `json:"product_name,omitempty"`

	// 厂商ID，未填写产品ID时厂商ID和型号必填
	ManufacturerId *string `json:"manufacturer_id,omitempty"`

	// 型号，未填写产品ID时厂商ID和型号必填
	Model *string `json:"model,omitempty"`

	// 产品的协议类型：0-mqtt，1-coap，2-modbus，3-http, 4-opcua
	ProtocolType *ProductRefererProtocolType `json:"protocol_type,omitempty"`

	// 产品类型：0-普通产品 1-网关产品
	ProductType *ProductRefererProductType `json:"product_type,omitempty"`
}

func (o ProductReferer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductReferer struct{}"
	}

	return strings.Join([]string{"ProductReferer", string(data)}, " ")
}

type ProductRefererProtocolType struct {
	value int32
}

type ProductRefererProtocolTypeEnum struct {
	E_0 ProductRefererProtocolType
	E_1 ProductRefererProtocolType
	E_2 ProductRefererProtocolType
	E_3 ProductRefererProtocolType
	E_4 ProductRefererProtocolType
}

func GetProductRefererProtocolTypeEnum() ProductRefererProtocolTypeEnum {
	return ProductRefererProtocolTypeEnum{
		E_0: ProductRefererProtocolType{
			value: 0,
		}, E_1: ProductRefererProtocolType{
			value: 1,
		}, E_2: ProductRefererProtocolType{
			value: 2,
		}, E_3: ProductRefererProtocolType{
			value: 3,
		}, E_4: ProductRefererProtocolType{
			value: 4,
		},
	}
}

func (c ProductRefererProtocolType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProductRefererProtocolType) UnmarshalJSON(b []byte) error {
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

type ProductRefererProductType struct {
	value int32
}

type ProductRefererProductTypeEnum struct {
	E_0 ProductRefererProductType
	E_1 ProductRefererProductType
}

func GetProductRefererProductTypeEnum() ProductRefererProductTypeEnum {
	return ProductRefererProductTypeEnum{
		E_0: ProductRefererProductType{
			value: 0,
		}, E_1: ProductRefererProductType{
			value: 1,
		},
	}
}

func (c ProductRefererProductType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProductRefererProductType) UnmarshalJSON(b []byte) error {
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
