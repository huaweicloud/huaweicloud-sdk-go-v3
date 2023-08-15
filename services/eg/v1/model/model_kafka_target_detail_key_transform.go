package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// KafkaTargetDetailKeyTransform 键值规则
type KafkaTargetDetailKeyTransform struct {

	// 键值规则类型
	Type *KafkaTargetDetailKeyTransformType `json:"type,omitempty"`

	// 键值规则，键值规则为VARIABLE，CONSTANT时必填
	Value *string `json:"value,omitempty"`

	// 键值规则模板，键值规则为VARIABLE时必填
	Template *string `json:"template,omitempty"`
}

func (o KafkaTargetDetailKeyTransform) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KafkaTargetDetailKeyTransform struct{}"
	}

	return strings.Join([]string{"KafkaTargetDetailKeyTransform", string(data)}, " ")
}

type KafkaTargetDetailKeyTransformType struct {
	value string
}

type KafkaTargetDetailKeyTransformTypeEnum struct {
	NONE     KafkaTargetDetailKeyTransformType
	VARIABLE KafkaTargetDetailKeyTransformType
	CONSTANT KafkaTargetDetailKeyTransformType
}

func GetKafkaTargetDetailKeyTransformTypeEnum() KafkaTargetDetailKeyTransformTypeEnum {
	return KafkaTargetDetailKeyTransformTypeEnum{
		NONE: KafkaTargetDetailKeyTransformType{
			value: "NONE",
		},
		VARIABLE: KafkaTargetDetailKeyTransformType{
			value: "VARIABLE",
		},
		CONSTANT: KafkaTargetDetailKeyTransformType{
			value: "CONSTANT",
		},
	}
}

func (c KafkaTargetDetailKeyTransformType) Value() string {
	return c.value
}

func (c KafkaTargetDetailKeyTransformType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KafkaTargetDetailKeyTransformType) UnmarshalJSON(b []byte) error {
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
