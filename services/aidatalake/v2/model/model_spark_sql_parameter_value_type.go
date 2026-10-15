package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SparkSqlParameterValueType **参数解释**：占位符类型，用于标识占位符值的数据类型。 **约束限制**：不涉及。 **取值范围**：   - STRING：字符串类型。   - DECIMAL：定点数类型。   - INTEGER：整型类型。   - DATE：日期类型。   - TIMESTAMP：时间类型。 **默认取值**：STRING。
type SparkSqlParameterValueType struct {
	value string
}

type SparkSqlParameterValueTypeEnum struct {
	STRING    SparkSqlParameterValueType
	DECIMAL   SparkSqlParameterValueType
	INTEGER   SparkSqlParameterValueType
	DATE      SparkSqlParameterValueType
	TIMESTAMP SparkSqlParameterValueType
}

func GetSparkSqlParameterValueTypeEnum() SparkSqlParameterValueTypeEnum {
	return SparkSqlParameterValueTypeEnum{
		STRING: SparkSqlParameterValueType{
			value: "STRING",
		},
		DECIMAL: SparkSqlParameterValueType{
			value: "DECIMAL",
		},
		INTEGER: SparkSqlParameterValueType{
			value: "INTEGER",
		},
		DATE: SparkSqlParameterValueType{
			value: "DATE",
		},
		TIMESTAMP: SparkSqlParameterValueType{
			value: "TIMESTAMP",
		},
	}
}

func (c SparkSqlParameterValueType) Value() string {
	return c.value
}

func (c SparkSqlParameterValueType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SparkSqlParameterValueType) UnmarshalJSON(b []byte) error {
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
