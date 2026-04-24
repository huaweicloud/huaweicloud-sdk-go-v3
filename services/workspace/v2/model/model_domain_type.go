package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DomainType 域类型 KERBEROS：对接AD OPEN_API：使用OPEN_API的方式对接统信域控
type DomainType struct {
	value string
}

type DomainTypeEnum struct {
	KERBEROS DomainType
	OPEN_API DomainType
}

func GetDomainTypeEnum() DomainTypeEnum {
	return DomainTypeEnum{
		KERBEROS: DomainType{
			value: "KERBEROS",
		},
		OPEN_API: DomainType{
			value: "OPEN_API",
		},
	}
}

func (c DomainType) Value() string {
	return c.value
}

func (c DomainType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DomainType) UnmarshalJSON(b []byte) error {
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
