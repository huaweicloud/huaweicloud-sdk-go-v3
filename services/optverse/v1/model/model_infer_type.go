package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InferType **参数解释**： 推理类型。 **约束限制**： 不涉及 **取值范围**： * online：在线推理 * edge：边缘推理 **默认取值**： 不涉及
type InferType struct {
	value string
}

type InferTypeEnum struct {
	ONLINE InferType
	EDGE   InferType
}

func GetInferTypeEnum() InferTypeEnum {
	return InferTypeEnum{
		ONLINE: InferType{
			value: "online",
		},
		EDGE: InferType{
			value: "edge",
		},
	}
}

func (c InferType) Value() string {
	return c.value
}

func (c InferType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InferType) UnmarshalJSON(b []byte) error {
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
