package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EndpointType 终端节点类型，取值范围: - EIP：本账号中的弹性公网IP - ECS：本账号中私网ECS实例 - ELB：本账号中私网ELB实例 - CUSTOM_IP：云外公网IP - CUSTOM_DOMAIN_NAME：云外公网域名 - CUSTOM_EIP：本Region的弹性公网IP
type EndpointType struct {
	value string
}

type EndpointTypeEnum struct {
	EIP                EndpointType
	ECS                EndpointType
	ELB                EndpointType
	CUSTOM_IP          EndpointType
	CUSTOM_DOMAIN_NAME EndpointType
	CUSTOM_EIP         EndpointType
}

func GetEndpointTypeEnum() EndpointTypeEnum {
	return EndpointTypeEnum{
		EIP: EndpointType{
			value: "EIP",
		},
		ECS: EndpointType{
			value: "ECS",
		},
		ELB: EndpointType{
			value: "ELB",
		},
		CUSTOM_IP: EndpointType{
			value: "CUSTOM_IP",
		},
		CUSTOM_DOMAIN_NAME: EndpointType{
			value: "CUSTOM_DOMAIN_NAME",
		},
		CUSTOM_EIP: EndpointType{
			value: "CUSTOM_EIP",
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
