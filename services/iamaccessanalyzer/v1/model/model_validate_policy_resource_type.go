package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ValidatePolicyResourceType 要附加到资源策略的资源类型。
type ValidatePolicyResourceType struct {
	value string
}

type ValidatePolicyResourceTypeEnum struct {
	IAMAGENCY ValidatePolicyResourceType
}

func GetValidatePolicyResourceTypeEnum() ValidatePolicyResourceTypeEnum {
	return ValidatePolicyResourceTypeEnum{
		IAMAGENCY: ValidatePolicyResourceType{
			value: "iam:agency",
		},
	}
}

func (c ValidatePolicyResourceType) Value() string {
	return c.value
}

func (c ValidatePolicyResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ValidatePolicyResourceType) UnmarshalJSON(b []byte) error {
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
