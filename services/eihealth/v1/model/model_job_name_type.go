package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobNameType 作业的名称类型
type JobNameType struct {
	value string
}

type JobNameTypeEnum struct {
	DATABASE JobNameType
	MANUAL   JobNameType
	AUTO     JobNameType
}

func GetJobNameTypeEnum() JobNameTypeEnum {
	return JobNameTypeEnum{
		DATABASE: JobNameType{
			value: "DATABASE",
		},
		MANUAL: JobNameType{
			value: "MANUAL",
		},
		AUTO: JobNameType{
			value: "AUTO",
		},
	}
}

func (c JobNameType) Value() string {
	return c.value
}

func (c JobNameType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobNameType) UnmarshalJSON(b []byte) error {
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
