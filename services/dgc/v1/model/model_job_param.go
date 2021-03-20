package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type JobParam struct {
	Name *string `json:"name,omitempty"`

	Value *string `json:"value,omitempty"`

	ParamType *JobParamParamType `json:"paramType,omitempty"`
}

func (o JobParam) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "JobParam struct{}"
	}

	return strings.Join([]string{"JobParam", string(data)}, " ")
}

type JobParamParamType struct {
	value string
}

type JobParamParamTypeEnum struct {
	VARIABLE  JobParamParamType
	CONSTANTS JobParamParamType
}

func GetJobParamParamTypeEnum() JobParamParamTypeEnum {
	return JobParamParamTypeEnum{
		VARIABLE: JobParamParamType{
			value: "variable",
		},
		CONSTANTS: JobParamParamType{
			value: "constants",
		},
	}
}

func (c JobParamParamType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *JobParamParamType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
