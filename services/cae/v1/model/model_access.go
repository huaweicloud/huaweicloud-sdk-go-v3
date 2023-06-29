package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Access struct {

	// 地址。
	Address *string `json:"address,omitempty"`

	// 类型。
	Type *AccessType `json:"type,omitempty"`
}

func (o Access) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Access struct{}"
	}

	return strings.Join([]string{"Access", string(data)}, " ")
}

type AccessType struct {
	value string
}

type AccessTypeEnum struct {
	LOAD_BALANCER AccessType
	INGRESS       AccessType
}

func GetAccessTypeEnum() AccessTypeEnum {
	return AccessTypeEnum{
		LOAD_BALANCER: AccessType{
			value: "LoadBalancer",
		},
		INGRESS: AccessType{
			value: "Ingress",
		},
	}
}

func (c AccessType) Value() string {
	return c.value
}

func (c AccessType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessType) UnmarshalJSON(b []byte) error {
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
