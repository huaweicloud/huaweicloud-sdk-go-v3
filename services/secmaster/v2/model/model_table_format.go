package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TableFormat **参数解释**: 表格式 - JSON Json格式 - DEBEZIUM_JSON Debezium JSON 格式 - CSV CSV格式 - PARQUET PARQUET格式 - ORC ORC格式  **约束限制** 不涉及 **取值范围**: - JSON - DEBEZIUM_JSON - CSV - PARQUET - ORC  **默认值** 不涉及
type TableFormat struct {
	value string
}

type TableFormatEnum struct {
	JSON          TableFormat
	DEBEZIUM_JSON TableFormat
	CSV           TableFormat
	PARQUET       TableFormat
	ORC           TableFormat
}

func GetTableFormatEnum() TableFormatEnum {
	return TableFormatEnum{
		JSON: TableFormat{
			value: "JSON",
		},
		DEBEZIUM_JSON: TableFormat{
			value: "DEBEZIUM_JSON",
		},
		CSV: TableFormat{
			value: "CSV",
		},
		PARQUET: TableFormat{
			value: "PARQUET",
		},
		ORC: TableFormat{
			value: "ORC",
		},
	}
}

func (c TableFormat) Value() string {
	return c.value
}

func (c TableFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableFormat) UnmarshalJSON(b []byte) error {
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
