package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// IsapTableColumnDto 表格列数据传输对象
type IsapTableColumnDto struct {

	// 表格列名称
	ColumnName *string `json:"column_name,omitempty"`

	// **参数解释**: 列字段类型 - PHYSICAL 物理字段 - METADATA 元数据字段 - VIRTUAL_METADATA 虚拟元数据字段 - COMPUTED 计算字段  **约束限制** 不涉及  **取值范围**: - PHYSICAL - METADATA - VIRTUAL_METADATA - COMPUTED  **默认值** 不涉及
	ColumnType *IsapTableColumnDtoColumnType `json:"column_type,omitempty"`

	// 表格列类型设置
	ColumnTypeSetting *string `json:"column_type_setting,omitempty"`

	// **参数解释**: 列字段数据类型 - ROW 行类型 - MAP_STRING 字符串映射类型 - MAP_DECIMAL 小数映射类型 - TINYINT 微整型 - SMALLINT 小整型 - INT 整型 - BIGINT 长整型 - DECIMAL 精确小数类型 - FLOAT 单精度浮点数 - DOUBLE 双精度浮点数 - CHAR 定长字符串 - VARCHAR 不定长字符串 - STRING 字符串类型 - KEYWORD 关键字类型 - BOOLEAN 布尔类型 - DATE 日期类型 - TIME 时间类型 - TIMESTAMP 时间戳类型 - TIMESTAMP_LTZ 本地时间戳类型  **约束限制** 不涉及  **取值范围**: - ROW - MAP_STRING - MAP_DECIMAL - TINYINT - SMALLINT - INT - BIGINT - DECIMAL - FLOAT - DOUBLE - CHAR - VARCHAR - STRING - KEYWORD - BOOLEAN - DATE - TIME - TIMESTAMP - TIMESTAMP_LTZ  **默认值** 不涉及
	ColumnDataType *IsapTableColumnDtoColumnDataType `json:"column_data_type,omitempty"`

	// 表格列数据类型设置
	ColumnDataTypeSetting *string `json:"column_data_type_setting,omitempty"`

	// 是否为空
	Nullable *bool `json:"nullable,omitempty"`

	// 是否为数组
	Array *bool `json:"array,omitempty"`

	ColumnDisplaySetting *ColumnDisplaySetting `json:"column_display_setting,omitempty"`

	// 列序号
	ColumnSequenceNumber *int32 `json:"column_sequence_number,omitempty"`
}

func (o IsapTableColumnDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapTableColumnDto struct{}"
	}

	return strings.Join([]string{"IsapTableColumnDto", string(data)}, " ")
}

type IsapTableColumnDtoColumnType struct {
	value string
}

type IsapTableColumnDtoColumnTypeEnum struct {
	PHYSICAL         IsapTableColumnDtoColumnType
	METADATA         IsapTableColumnDtoColumnType
	VIRTUAL_METADATA IsapTableColumnDtoColumnType
	COMPUTED         IsapTableColumnDtoColumnType
}

func GetIsapTableColumnDtoColumnTypeEnum() IsapTableColumnDtoColumnTypeEnum {
	return IsapTableColumnDtoColumnTypeEnum{
		PHYSICAL: IsapTableColumnDtoColumnType{
			value: "PHYSICAL",
		},
		METADATA: IsapTableColumnDtoColumnType{
			value: "METADATA",
		},
		VIRTUAL_METADATA: IsapTableColumnDtoColumnType{
			value: "VIRTUAL_METADATA",
		},
		COMPUTED: IsapTableColumnDtoColumnType{
			value: "COMPUTED",
		},
	}
}

func (c IsapTableColumnDtoColumnType) Value() string {
	return c.value
}

func (c IsapTableColumnDtoColumnType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IsapTableColumnDtoColumnType) UnmarshalJSON(b []byte) error {
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

type IsapTableColumnDtoColumnDataType struct {
	value string
}

type IsapTableColumnDtoColumnDataTypeEnum struct {
	ROW           IsapTableColumnDtoColumnDataType
	MAP_STRING    IsapTableColumnDtoColumnDataType
	MAP_DECIMAL   IsapTableColumnDtoColumnDataType
	TINYINT       IsapTableColumnDtoColumnDataType
	SMALLINT      IsapTableColumnDtoColumnDataType
	INT           IsapTableColumnDtoColumnDataType
	BIGINT        IsapTableColumnDtoColumnDataType
	DECIMAL       IsapTableColumnDtoColumnDataType
	FLOAT         IsapTableColumnDtoColumnDataType
	DOUBLE        IsapTableColumnDtoColumnDataType
	CHAR          IsapTableColumnDtoColumnDataType
	VARCHAR       IsapTableColumnDtoColumnDataType
	STRING        IsapTableColumnDtoColumnDataType
	KEYWORD       IsapTableColumnDtoColumnDataType
	BOOLEAN       IsapTableColumnDtoColumnDataType
	DATE          IsapTableColumnDtoColumnDataType
	TIME          IsapTableColumnDtoColumnDataType
	TIMESTAMP     IsapTableColumnDtoColumnDataType
	TIMESTAMP_LTZ IsapTableColumnDtoColumnDataType
}

func GetIsapTableColumnDtoColumnDataTypeEnum() IsapTableColumnDtoColumnDataTypeEnum {
	return IsapTableColumnDtoColumnDataTypeEnum{
		ROW: IsapTableColumnDtoColumnDataType{
			value: "ROW",
		},
		MAP_STRING: IsapTableColumnDtoColumnDataType{
			value: "MAP_STRING",
		},
		MAP_DECIMAL: IsapTableColumnDtoColumnDataType{
			value: "MAP_DECIMAL",
		},
		TINYINT: IsapTableColumnDtoColumnDataType{
			value: "TINYINT",
		},
		SMALLINT: IsapTableColumnDtoColumnDataType{
			value: "SMALLINT",
		},
		INT: IsapTableColumnDtoColumnDataType{
			value: "INT",
		},
		BIGINT: IsapTableColumnDtoColumnDataType{
			value: "BIGINT",
		},
		DECIMAL: IsapTableColumnDtoColumnDataType{
			value: "DECIMAL",
		},
		FLOAT: IsapTableColumnDtoColumnDataType{
			value: "FLOAT",
		},
		DOUBLE: IsapTableColumnDtoColumnDataType{
			value: "DOUBLE",
		},
		CHAR: IsapTableColumnDtoColumnDataType{
			value: "CHAR",
		},
		VARCHAR: IsapTableColumnDtoColumnDataType{
			value: "VARCHAR",
		},
		STRING: IsapTableColumnDtoColumnDataType{
			value: "STRING",
		},
		KEYWORD: IsapTableColumnDtoColumnDataType{
			value: "KEYWORD",
		},
		BOOLEAN: IsapTableColumnDtoColumnDataType{
			value: "BOOLEAN",
		},
		DATE: IsapTableColumnDtoColumnDataType{
			value: "DATE",
		},
		TIME: IsapTableColumnDtoColumnDataType{
			value: "TIME",
		},
		TIMESTAMP: IsapTableColumnDtoColumnDataType{
			value: "TIMESTAMP",
		},
		TIMESTAMP_LTZ: IsapTableColumnDtoColumnDataType{
			value: "TIMESTAMP_LTZ",
		},
	}
}

func (c IsapTableColumnDtoColumnDataType) Value() string {
	return c.value
}

func (c IsapTableColumnDtoColumnDataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IsapTableColumnDtoColumnDataType) UnmarshalJSON(b []byte) error {
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
