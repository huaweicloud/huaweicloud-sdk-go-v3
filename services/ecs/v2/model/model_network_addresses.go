package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NetworkAddresses struct {

	// IP地址。
	Addr string `json:"addr"`

	// IP地址类型，值为4或6。  4：IP地址类型是IPv4 6：IP地址类型是IPv6
	Version NetworkAddressesVersion `json:"version"`

	// IP地址对应的端口ID。
	OSEXTIPSportId string `json:"OS-EXT-IPS:port_id"`

	// 扩展属性，MAC地址。
	OSEXTIPSMACmacAddr string `json:"OS-EXT-IPS-MAC:mac_addr"`

	// 扩展属性，分配IP地址方式。  fixed：代表私有IP地址。 floating：代表浮动IP地址。
	OSEXTIPStype string `json:"OS-EXT-IPS:type"`

	// 是否是主网卡。  true：主网卡。 false：辅助网卡。
	Primary *bool `json:"primary,omitempty"`
}

func (o NetworkAddresses) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkAddresses struct{}"
	}

	return strings.Join([]string{"NetworkAddresses", string(data)}, " ")
}

type NetworkAddressesVersion struct {
	value int32
}

type NetworkAddressesVersionEnum struct {
	E_4 NetworkAddressesVersion
	E_6 NetworkAddressesVersion
}

func GetNetworkAddressesVersionEnum() NetworkAddressesVersionEnum {
	return NetworkAddressesVersionEnum{
		E_4: NetworkAddressesVersion{
			value: 4,
		}, E_6: NetworkAddressesVersion{
			value: 6,
		},
	}
}

func (c NetworkAddressesVersion) Value() int32 {
	return c.value
}

func (c NetworkAddressesVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NetworkAddressesVersion) UnmarshalJSON(b []byte) error {
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
