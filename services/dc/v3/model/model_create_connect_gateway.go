package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateConnectGateway 互联网关的实例对象
type CreateConnectGateway struct {

	// 网关名字
	Name string `json:"name"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 地址族信息 不填默认ipv4
	AddressFamily *CreateConnectGatewayAddressFamily `json:"address_family,omitempty"`
}

func (o CreateConnectGateway) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectGateway struct{}"
	}

	return strings.Join([]string{"CreateConnectGateway", string(data)}, " ")
}

type CreateConnectGatewayAddressFamily struct {
	value string
}

type CreateConnectGatewayAddressFamilyEnum struct {
	IPV4 CreateConnectGatewayAddressFamily
	DUAL CreateConnectGatewayAddressFamily
}

func GetCreateConnectGatewayAddressFamilyEnum() CreateConnectGatewayAddressFamilyEnum {
	return CreateConnectGatewayAddressFamilyEnum{
		IPV4: CreateConnectGatewayAddressFamily{
			value: "ipv4",
		},
		DUAL: CreateConnectGatewayAddressFamily{
			value: "dual",
		},
	}
}

func (c CreateConnectGatewayAddressFamily) Value() string {
	return c.value
}

func (c CreateConnectGatewayAddressFamily) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateConnectGatewayAddressFamily) UnmarshalJSON(b []byte) error {
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
