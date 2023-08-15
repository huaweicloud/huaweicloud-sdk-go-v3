package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddressInfo address数据结构说明
type AddressInfo struct {

	// IP地址版本。4：代表IPv4。6：代表IPv6。
	Version string `json:"version"`

	// IP地址
	Addr string `json:"addr"`

	// IP地址类型。fixed：代表私有IP地址。floating：代表浮动IP地址。
	OSEXTIPStype *AddressInfoOSEXTIPStype `json:"OS-EXT-IPS:type,omitempty"`

	// MAC地址。
	OSEXTIPSMACmacAddr *string `json:"OS-EXT-IPS-MAC:mac_addr,omitempty"`

	// IP地址对应的端口ID
	OSEXTIPSportId *string `json:"OS-EXT-IPS:port_id,omitempty"`
}

func (o AddressInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressInfo struct{}"
	}

	return strings.Join([]string{"AddressInfo", string(data)}, " ")
}

type AddressInfoOSEXTIPStype struct {
	value string
}

type AddressInfoOSEXTIPStypeEnum struct {
	FIXED    AddressInfoOSEXTIPStype
	FLOATING AddressInfoOSEXTIPStype
}

func GetAddressInfoOSEXTIPStypeEnum() AddressInfoOSEXTIPStypeEnum {
	return AddressInfoOSEXTIPStypeEnum{
		FIXED: AddressInfoOSEXTIPStype{
			value: "fixed",
		},
		FLOATING: AddressInfoOSEXTIPStype{
			value: "floating",
		},
	}
}

func (c AddressInfoOSEXTIPStype) Value() string {
	return c.value
}

func (c AddressInfoOSEXTIPStype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddressInfoOSEXTIPStype) UnmarshalJSON(b []byte) error {
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
