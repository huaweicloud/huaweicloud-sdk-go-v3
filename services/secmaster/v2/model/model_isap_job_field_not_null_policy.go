package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// IsapJobFieldNotNullPolicy 参数解释: 作业表非空字段处理策略  LOOSE 宽松 STRICT 严格 约束限制 不涉及 取值范围:  LOOSE STRICT 默认值 LOOSE
type IsapJobFieldNotNullPolicy struct {
	value string
}

type IsapJobFieldNotNullPolicyEnum struct {
	LOOSE  IsapJobFieldNotNullPolicy
	STRICT IsapJobFieldNotNullPolicy
}

func GetIsapJobFieldNotNullPolicyEnum() IsapJobFieldNotNullPolicyEnum {
	return IsapJobFieldNotNullPolicyEnum{
		LOOSE: IsapJobFieldNotNullPolicy{
			value: "LOOSE",
		},
		STRICT: IsapJobFieldNotNullPolicy{
			value: "STRICT",
		},
	}
}

func (c IsapJobFieldNotNullPolicy) Value() string {
	return c.value
}

func (c IsapJobFieldNotNullPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IsapJobFieldNotNullPolicy) UnmarshalJSON(b []byte) error {
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
