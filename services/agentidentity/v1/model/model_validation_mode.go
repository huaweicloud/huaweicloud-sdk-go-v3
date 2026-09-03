package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ValidationMode The validation mode for policy creation/updates. FAIL_ON_ANY_FINDINGS (default) fails if the Cedar analyzer reports any findings; IGNORE_ALL_FINDINGS allows creation/update even if findings are detected.
type ValidationMode struct {
	value string
}

type ValidationModeEnum struct {
	FAIL_ON_ANY_FINDINGS ValidationMode
	IGNORE_ALL_FINDINGS  ValidationMode
}

func GetValidationModeEnum() ValidationModeEnum {
	return ValidationModeEnum{
		FAIL_ON_ANY_FINDINGS: ValidationMode{
			value: "FAIL_ON_ANY_FINDINGS",
		},
		IGNORE_ALL_FINDINGS: ValidationMode{
			value: "IGNORE_ALL_FINDINGS",
		},
	}
}

func (c ValidationMode) Value() string {
	return c.value
}

func (c ValidationMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ValidationMode) UnmarshalJSON(b []byte) error {
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
