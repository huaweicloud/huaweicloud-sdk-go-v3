package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResourceType 资源的类型。
type ResourceType struct {
	value string
}

type ResourceTypeEnum struct {
	IAMAGENCY ResourceType
	IAMUSER   ResourceType
	KMSCMK    ResourceType
	OBSBUCKET ResourceType
}

func GetResourceTypeEnum() ResourceTypeEnum {
	return ResourceTypeEnum{
		IAMAGENCY: ResourceType{
			value: "iam:agency",
		},
		IAMUSER: ResourceType{
			value: "iam:user",
		},
		KMSCMK: ResourceType{
			value: "kms:cmk",
		},
		OBSBUCKET: ResourceType{
			value: "obs:bucket",
		},
	}
}

func (c ResourceType) Value() string {
	return c.value
}

func (c ResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceType) UnmarshalJSON(b []byte) error {
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
