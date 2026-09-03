package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VisibilityScopeEnum 可见性范围枚举。
type VisibilityScopeEnum struct {
	value string
}

type VisibilityScopeEnumEnum struct {
	OWNER_ONLY       VisibilityScopeEnum
	ALL_TENANTS      VisibilityScopeEnum
	SPECIFIC_TENANTS VisibilityScopeEnum
}

func GetVisibilityScopeEnumEnum() VisibilityScopeEnumEnum {
	return VisibilityScopeEnumEnum{
		OWNER_ONLY: VisibilityScopeEnum{
			value: "OWNER_ONLY",
		},
		ALL_TENANTS: VisibilityScopeEnum{
			value: "ALL_TENANTS",
		},
		SPECIFIC_TENANTS: VisibilityScopeEnum{
			value: "SPECIFIC_TENANTS",
		},
	}
}

func (c VisibilityScopeEnum) Value() string {
	return c.value
}

func (c VisibilityScopeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VisibilityScopeEnum) UnmarshalJSON(b []byte) error {
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
