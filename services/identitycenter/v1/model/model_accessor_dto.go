package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AccessorDto struct {
	Name string `json:"name"`

	Id string `json:"id"`

	AccessorType AccessorDtoAccessorType `json:"accessor_type"`
}

func (o AccessorDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessorDto struct{}"
	}

	return strings.Join([]string{"AccessorDto", string(data)}, " ")
}

type AccessorDtoAccessorType struct {
	value string
}

type AccessorDtoAccessorTypeEnum struct {
	GROUP AccessorDtoAccessorType
	USER  AccessorDtoAccessorType
}

func GetAccessorDtoAccessorTypeEnum() AccessorDtoAccessorTypeEnum {
	return AccessorDtoAccessorTypeEnum{
		GROUP: AccessorDtoAccessorType{
			value: "GROUP",
		},
		USER: AccessorDtoAccessorType{
			value: "USER",
		},
	}
}

func (c AccessorDtoAccessorType) Value() string {
	return c.value
}

func (c AccessorDtoAccessorType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessorDtoAccessorType) UnmarshalJSON(b []byte) error {
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
