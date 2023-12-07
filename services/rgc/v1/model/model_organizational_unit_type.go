package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OrganizationalUnitType OU类型。 * CORE - 核心OU * CUSTOM - 自定义OU * ROOT - 根OU
type OrganizationalUnitType struct {
	value string
}

type OrganizationalUnitTypeEnum struct {
	CORE   OrganizationalUnitType
	CUSTOM OrganizationalUnitType
	ROOT   OrganizationalUnitType
}

func GetOrganizationalUnitTypeEnum() OrganizationalUnitTypeEnum {
	return OrganizationalUnitTypeEnum{
		CORE: OrganizationalUnitType{
			value: "CORE",
		},
		CUSTOM: OrganizationalUnitType{
			value: "CUSTOM",
		},
		ROOT: OrganizationalUnitType{
			value: "ROOT",
		},
	}
}

func (c OrganizationalUnitType) Value() string {
	return c.value
}

func (c OrganizationalUnitType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OrganizationalUnitType) UnmarshalJSON(b []byte) error {
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
