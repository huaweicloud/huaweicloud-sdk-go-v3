package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DataTransformationProcessError **参数解释**: 数据加工处理错误 - NONE 无  **约束限制** 不涉及 **取值范围**: - NONE  **默认值** 不涉及
type DataTransformationProcessError struct {
	value string
}

type DataTransformationProcessErrorEnum struct {
	NONE DataTransformationProcessError
}

func GetDataTransformationProcessErrorEnum() DataTransformationProcessErrorEnum {
	return DataTransformationProcessErrorEnum{
		NONE: DataTransformationProcessError{
			value: "NONE",
		},
	}
}

func (c DataTransformationProcessError) Value() string {
	return c.value
}

func (c DataTransformationProcessError) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataTransformationProcessError) UnmarshalJSON(b []byte) error {
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
