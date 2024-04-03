package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddressSetDetailResponseDtoData 查询地址组详情数据
type AddressSetDetailResponseDtoData struct {

	// 地址组id
	Id *string `json:"id,omitempty"`

	// 地址组名称
	Name *string `json:"name,omitempty"`

	// 地址组描述
	Description *string `json:"description,omitempty"`

	// 地址组类型，0表示自定义地址组，1表示WAF回源IP地址组，2表示DDoS回源IP地址组，3表示NAT64转换地址组
	AddressSetType *int32 `json:"address_set_type,omitempty"`

	// 地址类型0 ipv4,1 ipv6
	AddressType *AddressSetDetailResponseDtoDataAddressType `json:"address_type,omitempty"`
}

func (o AddressSetDetailResponseDtoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressSetDetailResponseDtoData struct{}"
	}

	return strings.Join([]string{"AddressSetDetailResponseDtoData", string(data)}, " ")
}

type AddressSetDetailResponseDtoDataAddressType struct {
	value int32
}

type AddressSetDetailResponseDtoDataAddressTypeEnum struct {
	E_0 AddressSetDetailResponseDtoDataAddressType
	E_1 AddressSetDetailResponseDtoDataAddressType
}

func GetAddressSetDetailResponseDtoDataAddressTypeEnum() AddressSetDetailResponseDtoDataAddressTypeEnum {
	return AddressSetDetailResponseDtoDataAddressTypeEnum{
		E_0: AddressSetDetailResponseDtoDataAddressType{
			value: 0,
		}, E_1: AddressSetDetailResponseDtoDataAddressType{
			value: 1,
		},
	}
}

func (c AddressSetDetailResponseDtoDataAddressType) Value() int32 {
	return c.value
}

func (c AddressSetDetailResponseDtoDataAddressType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddressSetDetailResponseDtoDataAddressType) UnmarshalJSON(b []byte) error {
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
