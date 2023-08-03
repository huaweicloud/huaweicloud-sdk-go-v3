package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SmnTargetDetailSubjectTransform 标题规则
type SmnTargetDetailSubjectTransform struct {

	// 标题规则类型
	Type SmnTargetDetailSubjectTransformType `json:"type"`

	// 标题规则
	Value string `json:"value"`

	// 标题规则模板，键值规则为VARIABLE时必填
	Template *string `json:"template,omitempty"`
}

func (o SmnTargetDetailSubjectTransform) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmnTargetDetailSubjectTransform struct{}"
	}

	return strings.Join([]string{"SmnTargetDetailSubjectTransform", string(data)}, " ")
}

type SmnTargetDetailSubjectTransformType struct {
	value string
}

type SmnTargetDetailSubjectTransformTypeEnum struct {
	VARIABLE SmnTargetDetailSubjectTransformType
	CONSTANT SmnTargetDetailSubjectTransformType
}

func GetSmnTargetDetailSubjectTransformTypeEnum() SmnTargetDetailSubjectTransformTypeEnum {
	return SmnTargetDetailSubjectTransformTypeEnum{
		VARIABLE: SmnTargetDetailSubjectTransformType{
			value: "VARIABLE",
		},
		CONSTANT: SmnTargetDetailSubjectTransformType{
			value: "CONSTANT",
		},
	}
}

func (c SmnTargetDetailSubjectTransformType) Value() string {
	return c.value
}

func (c SmnTargetDetailSubjectTransformType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmnTargetDetailSubjectTransformType) UnmarshalJSON(b []byte) error {
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
