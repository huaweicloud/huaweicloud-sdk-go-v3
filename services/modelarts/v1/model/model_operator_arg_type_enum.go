package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OperatorArgTypeEnum 算子参数的类型枚举，表示参数值的数据类型，枚举值： - INT - FLOAT - STRING - ENUM - LIST - BOOLEAN - LARGE_MODEL - OBS - LABEL_FILTER - DATASET - DYNAMIC_INPUT_LIST - PROMPT - OUTPUT_TASK_CATEGORY
type OperatorArgTypeEnum struct {
	value string
}

type OperatorArgTypeEnumEnum struct {
	INT                  OperatorArgTypeEnum
	FLOAT                OperatorArgTypeEnum
	STRING               OperatorArgTypeEnum
	ENUM                 OperatorArgTypeEnum
	LIST                 OperatorArgTypeEnum
	BOOLEAN              OperatorArgTypeEnum
	LARGE_MODEL          OperatorArgTypeEnum
	OBS                  OperatorArgTypeEnum
	LABEL_FILTER         OperatorArgTypeEnum
	DATASET              OperatorArgTypeEnum
	DYNAMIC_INPUT_LIST   OperatorArgTypeEnum
	PROMPT               OperatorArgTypeEnum
	OUTPUT_TASK_CATEGORY OperatorArgTypeEnum
}

func GetOperatorArgTypeEnumEnum() OperatorArgTypeEnumEnum {
	return OperatorArgTypeEnumEnum{
		INT: OperatorArgTypeEnum{
			value: "INT",
		},
		FLOAT: OperatorArgTypeEnum{
			value: "FLOAT",
		},
		STRING: OperatorArgTypeEnum{
			value: "STRING",
		},
		ENUM: OperatorArgTypeEnum{
			value: "ENUM",
		},
		LIST: OperatorArgTypeEnum{
			value: "LIST",
		},
		BOOLEAN: OperatorArgTypeEnum{
			value: "BOOLEAN",
		},
		LARGE_MODEL: OperatorArgTypeEnum{
			value: "LARGE_MODEL",
		},
		OBS: OperatorArgTypeEnum{
			value: "OBS",
		},
		LABEL_FILTER: OperatorArgTypeEnum{
			value: "LABEL_FILTER",
		},
		DATASET: OperatorArgTypeEnum{
			value: "DATASET",
		},
		DYNAMIC_INPUT_LIST: OperatorArgTypeEnum{
			value: "DYNAMIC_INPUT_LIST",
		},
		PROMPT: OperatorArgTypeEnum{
			value: "PROMPT",
		},
		OUTPUT_TASK_CATEGORY: OperatorArgTypeEnum{
			value: "OUTPUT_TASK_CATEGORY",
		},
	}
}

func (c OperatorArgTypeEnum) Value() string {
	return c.value
}

func (c OperatorArgTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OperatorArgTypeEnum) UnmarshalJSON(b []byte) error {
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
