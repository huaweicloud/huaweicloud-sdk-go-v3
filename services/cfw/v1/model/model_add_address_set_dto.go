package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddAddressSetDto 添加地址组请求体
type AddAddressSetDto struct {

	// 互联网边界防护对象id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)，type为0的为互联网边界防护对象id。
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
