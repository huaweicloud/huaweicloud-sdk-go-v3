package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GatewayTypeEnum 网关的类型。GDGW：全球接入网关。
type GatewayTypeEnum struct {
	value string
}

type GatewayTypeEnumEnum struct {
	GDGW GatewayTypeEnum
}

func GetGatewayTypeEnumEnum() GatewayTypeEnumEnum {
	return GatewayTypeEnumEnum{
		GDGW: GatewayTypeEnum{
			value: "GDGW",
		},
	}
}

func (c GatewayTypeEnum) Value() string {
	return c.value
}

func (c GatewayTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GatewayTypeEnum) UnmarshalJSON(b []byte) error {
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
