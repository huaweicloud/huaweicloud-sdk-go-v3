package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DataMaskParas struct {

	// 敏感字段。
	ColumnName *string `json:"column_name,omitempty"`

	// 算法名称。
	AlgorithmName *string `json:"algorithm_name,omitempty"`

	// 算法类型。
	AlgorithmType *string `json:"algorithm_type,omitempty"`

	// 算法名称。
	EnName *string `json:"en_name,omitempty"`

	// 参数。
	AlgorithmParameters *string `json:"algorithm_parameters,omitempty"`

	FailurePolicy *DataMaskParasFailurePolicy `json:"failure_policy,omitempty"`
}

func (o DataMaskParas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataMaskParas struct{}"
	}

	return strings.Join([]string{"DataMaskParas", string(data)}, " ")
}

type DataMaskParasFailurePolicy struct {
	value string
}

type DataMaskParasFailurePolicyEnum struct {
	SKIP                    DataMaskParasFailurePolicy
	INTERRUPT_AND_EXCEPTION DataMaskParasFailurePolicy
	SET_NULL                DataMaskParasFailurePolicy
	DEFAULT_VALUE           DataMaskParasFailurePolicy
}

func GetDataMaskParasFailurePolicyEnum() DataMaskParasFailurePolicyEnum {
	return DataMaskParasFailurePolicyEnum{
		SKIP: DataMaskParasFailurePolicy{
			value: "SKIP",
		},
		INTERRUPT_AND_EXCEPTION: DataMaskParasFailurePolicy{
			value: "INTERRUPT_AND_EXCEPTION",
		},
		SET_NULL: DataMaskParasFailurePolicy{
			value: "SET_NULL",
		},
		DEFAULT_VALUE: DataMaskParasFailurePolicy{
			value: "DEFAULT_VALUE",
		},
	}
}

func (c DataMaskParasFailurePolicy) Value() string {
	return c.value
}

func (c DataMaskParasFailurePolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataMaskParasFailurePolicy) UnmarshalJSON(b []byte) error {
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
