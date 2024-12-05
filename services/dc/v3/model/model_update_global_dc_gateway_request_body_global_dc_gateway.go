package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateGlobalDcGatewayRequestBodyGlobalDcGateway struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 地址簇类型
	AddressFamily *UpdateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamily `json:"address_family,omitempty"`
}

func (o UpdateGlobalDcGatewayRequestBodyGlobalDcGateway) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalDcGatewayRequestBodyGlobalDcGateway struct{}"
	}

	return strings.Join([]string{"UpdateGlobalDcGatewayRequestBodyGlobalDcGateway", string(data)}, " ")
}

type UpdateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamily struct {
	value string
}

type UpdateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamilyEnum struct {
	IPV4 UpdateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamily
	DUAL UpdateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamily
}

func GetUpdateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamilyEnum() UpdateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamilyEnum {
	return UpdateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamilyEnum{
		IPV4: UpdateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamily{
			value: "ipv4",
		},
		DUAL: UpdateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamily{
			value: "dual",
		},
	}
}

func (c UpdateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamily) Value() string {
	return c.value
}

func (c UpdateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamily) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamily) UnmarshalJSON(b []byte) error {
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
