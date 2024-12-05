package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateConnectGateway struct {

	// 网关名字
	Name *string `json:"name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 地址族信息
	AddressFamily *UpdateConnectGatewayAddressFamily `json:"address_family,omitempty"`
}

func (o UpdateConnectGateway) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConnectGateway struct{}"
	}

	return strings.Join([]string{"UpdateConnectGateway", string(data)}, " ")
}

type UpdateConnectGatewayAddressFamily struct {
	value string
}

type UpdateConnectGatewayAddressFamilyEnum struct {
	IPV4 UpdateConnectGatewayAddressFamily
	DUAL UpdateConnectGatewayAddressFamily
}

func GetUpdateConnectGatewayAddressFamilyEnum() UpdateConnectGatewayAddressFamilyEnum {
	return UpdateConnectGatewayAddressFamilyEnum{
		IPV4: UpdateConnectGatewayAddressFamily{
			value: "ipv4",
		},
		DUAL: UpdateConnectGatewayAddressFamily{
			value: "dual",
		},
	}
}

func (c UpdateConnectGatewayAddressFamily) Value() string {
	return c.value
}

func (c UpdateConnectGatewayAddressFamily) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateConnectGatewayAddressFamily) UnmarshalJSON(b []byte) error {
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
