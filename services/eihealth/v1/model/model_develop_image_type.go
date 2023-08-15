package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DevelopImageType struct {
	value string
}

type DevelopImageTypeEnum struct {
	SYSTEM   DevelopImageType
	CUSTOMER DevelopImageType
}

func GetDevelopImageTypeEnum() DevelopImageTypeEnum {
	return DevelopImageTypeEnum{
		SYSTEM: DevelopImageType{
			value: "SYSTEM",
		},
		CUSTOMER: DevelopImageType{
			value: "CUSTOMER",
		},
	}
}

func (c DevelopImageType) Value() string {
	return c.value
}

func (c DevelopImageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DevelopImageType) UnmarshalJSON(b []byte) error {
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
