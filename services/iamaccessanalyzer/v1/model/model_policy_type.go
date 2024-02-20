package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PolicyType 要校验的策略类型。
type PolicyType struct {
	value string
}

type PolicyTypeEnum struct {
	IDENTITY_POLICY        PolicyType
	RESOURCE_POLICY        PolicyType
	SERVICE_CONTROL_POLICY PolicyType
}

func GetPolicyTypeEnum() PolicyTypeEnum {
	return PolicyTypeEnum{
		IDENTITY_POLICY: PolicyType{
			value: "identity_policy",
		},
		RESOURCE_POLICY: PolicyType{
			value: "resource_policy",
		},
		SERVICE_CONTROL_POLICY: PolicyType{
			value: "service_control_policy",
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
