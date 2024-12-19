package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddressFamily 地址族类型: - ipv4: ipv4 - ipv6: ipv6
type AddressFamily struct {
	value string
}

type AddressFamilyEnum struct {
	IPV4 AddressFamily
	IPV6 AddressFamily
}

func GetAddressFamilyEnum() AddressFamilyEnum {
	return AddressFamilyEnum{
		IPV4: AddressFamily{
			value: "ipv4",
		},
		IPV6: AddressFamily{
			value: "ipv6",
		},
	}
}

func (c AddressFamily) Value() string {
	return c.value
}

func (c AddressFamily) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddressFamily) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
