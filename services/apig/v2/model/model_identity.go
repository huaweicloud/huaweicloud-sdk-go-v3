package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Identity struct {

	// 参数名称
	Name string `json:"name"`

	// 参数位置
	Location IdentityLocation `json:"location"`

	// 参数校验表达式，默认为null，不做校验
	Validation *string `json:"validation,omitempty"`
}

func (o Identity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Identity struct{}"
	}

	return strings.Join([]string{"Identity", string(data)}, " ")
}

type IdentityLocation struct {
	value string
}

type IdentityLocationEnum struct {
	HEADER IdentityLocation
	QUERY  IdentityLocation
}

func GetIdentityLocationEnum() IdentityLocationEnum {
	return IdentityLocationEnum{
		HEADER: IdentityLocation{
			value: "HEADER",
		},
		QUERY: IdentityLocation{
			value: "QUERY",
		},
	}
}

func (c IdentityLocation) Value() string {
	return c.value
}

func (c IdentityLocation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IdentityLocation) UnmarshalJSON(b []byte) error {
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
