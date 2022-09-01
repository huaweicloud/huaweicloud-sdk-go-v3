package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListProductsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 每页显示条目数量，最大数量999，超过999后只返回999
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 应用ID
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	// 产品ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 产品名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 厂商ID
	ManufacturerId *string `json:"manufacturer_id,omitempty" xml:"manufacturer_id"`

	// 厂商名称
	ManufacturerName *string `json:"manufacturer_name,omitempty" xml:"manufacturer_name"`

	// 型号
	Model *string `json:"model,omitempty" xml:"model"`

	// 产品的设备类型，默认Default
	DeviceType *string `json:"device_type,omitempty" xml:"device_type"`

	// 产品类型，0-普通产品(不支持子设备) 1-网关产品
	ProductType *ListProductsRequestProductType `json:"product_type,omitempty" xml:"product_type"`

	// 产品的协议类型 0-mqtt 2-modbus 4-opcua
	ProtocolType *ListProductsRequestProtocolType `json:"protocol_type,omitempty" xml:"protocol_type"`

	// 创建用户
	CreatedUserName *string `json:"created_user_name,omitempty" xml:"created_user_name"`

	// 创建时间起始，格式timestamp(ms)，使用UTC时区
	CreatedDateStart *int64 `json:"created_date_start,omitempty" xml:"created_date_start"`

	// 创建时间截止，格式timestamp(ms)，使用UTC时区
	CreatedDateEnd *int64 `json:"created_date_end,omitempty" xml:"created_date_end"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 应用名称
	AppName *string `json:"app_name,omitempty" xml:"app_name"`

	// 产品唯一序列（系统唯一值，用于MQS的TOPIC中标记产品）
	ProductSerial *string `json:"product_serial,omitempty" xml:"product_serial"`
}

func (o ListProductsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsRequest struct{}"
	}

	return strings.Join([]string{"ListProductsRequest", string(data)}, " ")
}

type ListProductsRequestProductType struct {
	value int32
}

type ListProductsRequestProductTypeEnum struct {
	E_0 ListProductsRequestProductType
	E_1 ListProductsRequestProductType
}

func GetListProductsRequestProductTypeEnum() ListProductsRequestProductTypeEnum {
	return ListProductsRequestProductTypeEnum{
		E_0: ListProductsRequestProductType{
			value: 0,
		}, E_1: ListProductsRequestProductType{
			value: 1,
		},
	}
}

func (c ListProductsRequestProductType) Value() int32 {
	return c.value
}

func (c ListProductsRequestProductType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProductsRequestProductType) UnmarshalJSON(b []byte) error {
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

type ListProductsRequestProtocolType struct {
	value int32
}

type ListProductsRequestProtocolTypeEnum struct {
	E_0 ListProductsRequestProtocolType
	E_2 ListProductsRequestProtocolType
	E_4 ListProductsRequestProtocolType
}

func GetListProductsRequestProtocolTypeEnum() ListProductsRequestProtocolTypeEnum {
	return ListProductsRequestProtocolTypeEnum{
		E_0: ListProductsRequestProtocolType{
			value: 0,
		}, E_2: ListProductsRequestProtocolType{
			value: 2,
		}, E_4: ListProductsRequestProtocolType{
			value: 4,
		},
	}
}

func (c ListProductsRequestProtocolType) Value() int32 {
	return c.value
}

func (c ListProductsRequestProtocolType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProductsRequestProtocolType) UnmarshalJSON(b []byte) error {
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
