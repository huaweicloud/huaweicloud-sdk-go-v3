package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobParam 作业参数
type JobParam struct {

	// 参数名称，只能包含英文字母、数字、中划线和下划线。
	Name string `json:"name"`

	// 参数值
	Value string `json:"value"`

	// 参数类型:   - variable: 变量   - constants: 常量
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
	VARIABLE  JobParamType
	CONSTANTS JobParamType
}

func GetJobParamTypeEnum() JobParamTypeEnum {
	return JobParamTypeEnum{
		VARIABLE: JobParamType{
			value: "variable",
		},
		CONSTANTS: JobParamType{
			value: "constants",
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
