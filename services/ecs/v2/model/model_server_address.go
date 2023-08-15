package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ServerAddress 弹性云服务器的网络属性。
type ServerAddress struct {

	// IP地址版本。  - “4”：代表IPv4。 - “6”：代表IPv6。
	Version string `json:"version"`

	// IP地址。
	Addr string `json:"addr"`

	// IP地址类型。  - fixed：代表私有IP地址。 - floating：代表浮动IP地址。
	OSEXTIPStype *ServerAddressOSEXTIPStype `json:"OS-EXT-IPS:type,omitempty"`

	// MAC地址。
	OSEXTIPSMACmacAddr *string `json:"OS-EXT-IPS-MAC:mac_addr,omitempty"`

	// IP地址对应的端口ID。
	OSEXTIPSportId *string `json:"OS-EXT-IPS:port_id,omitempty"`
}

func (o ServerAddress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerAddress struct{}"
	}

	return strings.Join([]string{"ServerAddress", string(data)}, " ")
}

type ServerAddressOSEXTIPStype struct {
	value string
}

type ServerAddressOSEXTIPStypeEnum struct {
	FIXED    ServerAddressOSEXTIPStype
	FLOATING ServerAddressOSEXTIPStype
}

func GetServerAddressOSEXTIPStypeEnum() ServerAddressOSEXTIPStypeEnum {
	return ServerAddressOSEXTIPStypeEnum{
		FIXED: ServerAddressOSEXTIPStype{
			value: "fixed",
		},
		FLOATING: ServerAddressOSEXTIPStype{
			value: "floating",
		},
	}
}

func (c ServerAddressOSEXTIPStype) Value() string {
	return c.value
}

func (c ServerAddressOSEXTIPStype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerAddressOSEXTIPStype) UnmarshalJSON(b []byte) error {
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
