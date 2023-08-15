package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateAddressSetDto 更新地址组信息
type UpdateAddressSetDto struct {

	// 地址组名称
	Name *string `json:"name,omitempty"`

	// 地址组描述
	Description *string `json:"description,omitempty"`

	// 地址类型0 ipv4,1 ipv6,2 domain
	AddressType *UpdateAddressSetDtoAddressType `json:"address_type,omitempty"`
}

func (o UpdateAddressSetDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAddressSetDto struct{}"
	}

	return strings.Join([]string{"UpdateAddressSetDto", string(data)}, " ")
}

type UpdateAddressSetDtoAddressType struct {
	value int32
}

type UpdateAddressSetDtoAddressTypeEnum struct {
	E_0 UpdateAddressSetDtoAddressType
	E_1 UpdateAddressSetDtoAddressType
	E_2 UpdateAddressSetDtoAddressType
}

func GetUpdateAddressSetDtoAddressTypeEnum() UpdateAddressSetDtoAddressTypeEnum {
	return UpdateAddressSetDtoAddressTypeEnum{
		E_0: UpdateAddressSetDtoAddressType{
			value: 0,
		}, E_1: UpdateAddressSetDtoAddressType{
			value: 1,
		}, E_2: UpdateAddressSetDtoAddressType{
			value: 2,
		},
	}
}

func (c UpdateAddressSetDtoAddressType) Value() int32 {
	return c.value
}

func (c UpdateAddressSetDtoAddressType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAddressSetDtoAddressType) UnmarshalJSON(b []byte) error {
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
