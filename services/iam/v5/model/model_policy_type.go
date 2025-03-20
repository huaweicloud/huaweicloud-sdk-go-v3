package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PolicyType 身份策略类型，可以为“自定义”（custom）或“系统预置”（system）。
type PolicyType struct {
	value string
}

type PolicyTypeEnum struct {
	CUSTOM PolicyType
	SYSTEM PolicyType
}

func GetPolicyTypeEnum() PolicyTypeEnum {
	return PolicyTypeEnum{
		CUSTOM: PolicyType{
			value: "custom",
		},
		SYSTEM: PolicyType{
			value: "system",
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
