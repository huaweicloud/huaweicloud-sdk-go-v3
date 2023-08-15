package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ColumnStatisticsObj struct {

	// 数据列名
	ColumnName string `json:"column_name"`

	// 数据类型，字段类型包括array bigint binary boolean char date decimal double float int interval map set smallint string struct timestamp tinyint union varchar
	ColumnType string `json:"column_type"`

	// 统计信息类型
	DataType ColumnStatisticsObjDataType `json:"data_type"`

	BinaryStatisticsData *BinaryColumnStatisticsData `json:"binary_statistics_data,omitempty"`

	LongStatisticsData *LongColumnStatisticsData `json:"long_statistics_data,omitempty"`

	DecimalStatisticsData *DecimalColumnStatisticsData `json:"decimal_statistics_data,omitempty"`

	StringStatisticsData *StringColumnStatisticsData `json:"string_statistics_data,omitempty"`

	DoubleStatisticsData *DoubleColumnStatisticsData `json:"double_statistics_data,omitempty"`

	DateStatisticsData *DateColumnStatisticsData `json:"date_statistics_data,omitempty"`

	BooleanStatisticsData *BooleanColumnStatisticsData `json:"boolean_statistics_data,omitempty"`
}

func (o ColumnStatisticsObj) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColumnStatisticsObj struct{}"
	}

	return strings.Join([]string{"ColumnStatisticsObj", string(data)}, " ")
}

type ColumnStatisticsObjDataType struct {
	value string
}

type ColumnStatisticsObjDataTypeEnum struct {
	BINARY_STATS  ColumnStatisticsObjDataType
	BOOLEAN_STATS ColumnStatisticsObjDataType
	DATE_STATS    ColumnStatisticsObjDataType
	DECIMAL_STATS ColumnStatisticsObjDataType
	DOUBLE_STATS  ColumnStatisticsObjDataType
	LONG_STATS    ColumnStatisticsObjDataType
	STRING_STATS  ColumnStatisticsObjDataType
}

func GetColumnStatisticsObjDataTypeEnum() ColumnStatisticsObjDataTypeEnum {
	return ColumnStatisticsObjDataTypeEnum{
		BINARY_STATS: ColumnStatisticsObjDataType{
			value: "binaryStats",
		},
		BOOLEAN_STATS: ColumnStatisticsObjDataType{
			value: "booleanStats",
		},
		DATE_STATS: ColumnStatisticsObjDataType{
			value: "dateStats",
		},
		DECIMAL_STATS: ColumnStatisticsObjDataType{
			value: "decimalStats",
		},
		DOUBLE_STATS: ColumnStatisticsObjDataType{
			value: "doubleStats",
		},
		LONG_STATS: ColumnStatisticsObjDataType{
			value: "longStats",
		},
		STRING_STATS: ColumnStatisticsObjDataType{
			value: "stringStats",
		},
	}
}

func (c ColumnStatisticsObjDataType) Value() string {
	return c.value
}

func (c ColumnStatisticsObjDataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ColumnStatisticsObjDataType) UnmarshalJSON(b []byte) error {
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
