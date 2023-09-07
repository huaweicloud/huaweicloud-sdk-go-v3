package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type JobParam struct {
	Name string `json:"name"`

	Value string `json:"value"`

	// 参数类型
	Type *JobParamType `json:"type,omitempty"`
}

func (o JobParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobParam struct{}"
	}

	return strings.Join([]string{"JobParam", string(data)}, " ")
}

type JobParamType struct {
	value string
}

type JobParamTypeEnum struct {
	PROCEDURE JobParamType
	VARIABLE  JobParamType
}

func GetJobParamTypeEnum() JobParamTypeEnum {
	return JobParamTypeEnum{
		PROCEDURE: JobParamType{
			value: "procedure",
		},
		VARIABLE: JobParamType{
			value: "variable",
		},
	}
}

func (c JobParamType) Value() string {
	return c.value
}

func (c JobParamType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobParamType) UnmarshalJSON(b []byte) error {
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
