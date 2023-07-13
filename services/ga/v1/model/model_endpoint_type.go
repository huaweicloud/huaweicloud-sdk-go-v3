package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EndpointType 终端节点类型。
type EndpointType struct {
	value string
}

type EndpointTypeEnum struct {
	EIP EndpointType
}

func GetEndpointTypeEnum() EndpointTypeEnum {
	return EndpointTypeEnum{
		EIP: EndpointType{
			value: "EIP",
		},
	}
}

func (c EndpointType) Value() string {
	return c.value
}

func (c EndpointType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointType) UnmarshalJSON(b []byte) error {
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
