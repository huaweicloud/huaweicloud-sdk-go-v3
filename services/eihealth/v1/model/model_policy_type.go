package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PolicyType struct {
	value string
}

type PolicyTypeEnum struct {
	ALLOW PolicyType
	DENY  PolicyType
}

func GetPolicyTypeEnum() PolicyTypeEnum {
	return PolicyTypeEnum{
		ALLOW: PolicyType{
			value: "allow",
		},
		DENY: PolicyType{
			value: "deny",
		},
	}
}

func (c PolicyType) Value() string {
	return c.value
}

func (c PolicyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyType) UnmarshalJSON(b []byte) error {
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
