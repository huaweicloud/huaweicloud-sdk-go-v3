package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OrganizationalUnitTypeForSetup 可设置的OU类型。包括核心OU和自定义OU。
type OrganizationalUnitTypeForSetup struct {
	value string
}

type OrganizationalUnitTypeForSetupEnum struct {
	CORE   OrganizationalUnitTypeForSetup
	CUSTOM OrganizationalUnitTypeForSetup
}

func GetOrganizationalUnitTypeForSetupEnum() OrganizationalUnitTypeForSetupEnum {
	return OrganizationalUnitTypeForSetupEnum{
		CORE: OrganizationalUnitTypeForSetup{
			value: "CORE",
		},
		CUSTOM: OrganizationalUnitTypeForSetup{
			value: "CUSTOM",
		},
	}
}

func (c OrganizationalUnitTypeForSetup) Value() string {
	return c.value
}

func (c OrganizationalUnitTypeForSetup) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OrganizationalUnitTypeForSetup) UnmarshalJSON(b []byte) error {
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
