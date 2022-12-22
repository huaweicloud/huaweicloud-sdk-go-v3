package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 添加地址组请求体
type AddAddressSetDto struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用查询防火墙实例接口获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。具体可参考APIExlorer和帮助中心FAQ。
	ObjectId string `json:"object_id"`

	// 地址组名称
	Name string `json:"name"`

	// 地址组描述
	Description *string `json:"description,omitempty"`

	// 地址类型0 ipv4,1 ipv6
	AddressType *AddAddressSetDtoAddressType `json:"address_type,omitempty"`
}

func (o AddAddressSetDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAddressSetDto struct{}"
	}

	return strings.Join([]string{"AddAddressSetDto", string(data)}, " ")
}

type AddAddressSetDtoAddressType struct {
	value int32
}

type AddAddressSetDtoAddressTypeEnum struct {
	E_0 AddAddressSetDtoAddressType
	E_1 AddAddressSetDtoAddressType
}

func GetAddAddressSetDtoAddressTypeEnum() AddAddressSetDtoAddressTypeEnum {
	return AddAddressSetDtoAddressTypeEnum{
		E_0: AddAddressSetDtoAddressType{
			value: 0,
		}, E_1: AddAddressSetDtoAddressType{
			value: 1,
		},
	}
}

func (c AddAddressSetDtoAddressType) Value() int32 {
	return c.value
}

func (c AddAddressSetDtoAddressType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddAddressSetDtoAddressType) UnmarshalJSON(b []byte) error {
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
