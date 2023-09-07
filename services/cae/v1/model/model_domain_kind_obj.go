package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DomainKindObj API类型，固定值“Domain”，该值不可修改。
type DomainKindObj struct {
	value string
}

type DomainKindObjEnum struct {
	DOMAIN DomainKindObj
}

func GetDomainKindObjEnum() DomainKindObjEnum {
	return DomainKindObjEnum{
		DOMAIN: DomainKindObj{
			value: "Domain",
		},
	}
}

func (c DomainKindObj) Value() string {
	return c.value
}

func (c DomainKindObj) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DomainKindObj) UnmarshalJSON(b []byte) error {
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
