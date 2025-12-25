package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TableColumnForIsapPipe 管道对应表的列详情
type TableColumnForIsapPipe struct {

	// 表列名称
	ColumnName string `json:"column_name"`

	// **参数解释**: 列字段数据类型 - ROW 行类型 - MAP_STRING 字符串映射类型 - MAP_DECIMAL 小数映射类型 - TINYINT 微整型 - SMALLINT 小整型 - INT 整型 - BIGINT 长整型 - DECIMAL 精确小数类型 - FLOAT 单精度浮点数 - DOUBLE 双精度浮点数 - CHAR 定长字符串 - VARCHAR 不定长字符串 - STRING 字符串类型 - KEYWORD 关键字类型 - BOOLEAN 布尔类型 - DATE 日期类型 - TIME 时间类型 - TIMESTAMP 时间戳类型 - TIMESTAMP_LTZ 本地时间戳类型  **约束限制** 不涉及  **取值范围**: - ROW - MAP_STRING - MAP_DECIMAL - TINYINT - SMALLINT - INT - BIGINT - DECIMAL - FLOAT - DOUBLE - CHAR - VARCHAR - STRING - KEYWORD - BOOLEAN - DATE - TIME - TIMESTAMP - TIMESTAMP_LTZ  **默认值** 不涉及
	ColumnDataType TableColumnForIsapPipeColumnDataType `json:"column_data_type"`

	// 列数据类型设置
	ColumnDataTypeSetting *string `json:"column_data_type_setting,omitempty"`

	// **参数解释**: 列字段类型 - PHYSICAL 物理字段 - METADATA 元数据字段 - VIRTUAL_METADATA 虚拟元数据字段 - COMPUTED 计算字段  **约束限制** 不涉及  **取值范围**: - PHYSICAL - METADATA - VIRTUAL_METADATA - COMPUTED  **默认值** 不涉及
	ColumnType *TableColumnForIsapPipeColumnType `json:"column_type,omitempty"`

	// 列类型设置
	ColumnTypeSetting *string `json:"column_type_setting,omitempty"`

	// 列默认值
	DefaultValue *string `json:"default_value,omitempty"`

	// 列是否可为空
	Nullable *bool `json:"nullable,omitempty"`

	// 是否为数组
	Array *bool `json:"array,omitempty"`

	// 列深度
	Depth *int64 `json:"depth,omitempty"`

	// 列名称
	ParentName *string `json:"parent_name,omitempty"`

	// 列名称
	OwnName *string `json:"own_name,omitempty"`
}

func (o TableColumnForIsapPipe) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableColumnForIsapPipe struct{}"
	}

	return strings.Join([]string{"TableColumnForIsapPipe", string(data)}, " ")
}

type TableColumnForIsapPipeColumnDataType struct {
	value string
}

type TableColumnForIsapPipeColumnDataTypeEnum struct {
	ROW           TableColumnForIsapPipeColumnDataType
	MAP_STRING    TableColumnForIsapPipeColumnDataType
	MAP_DECIMAL   TableColumnForIsapPipeColumnDataType
	TINYINT       TableColumnForIsapPipeColumnDataType
	SMALLINT      TableColumnForIsapPipeColumnDataType
	INT           TableColumnForIsapPipeColumnDataType
	BIGINT        TableColumnForIsapPipeColumnDataType
	DECIMAL       TableColumnForIsapPipeColumnDataType
	FLOAT         TableColumnForIsapPipeColumnDataType
	DOUBLE        TableColumnForIsapPipeColumnDataType
	CHAR          TableColumnForIsapPipeColumnDataType
	VARCHAR       TableColumnForIsapPipeColumnDataType
	STRING        TableColumnForIsapPipeColumnDataType
	BOOLEAN       TableColumnForIsapPipeColumnDataType
	DATE          TableColumnForIsapPipeColumnDataType
	TIME          TableColumnForIsapPipeColumnDataType
	TIMESTAMP     TableColumnForIsapPipeColumnDataType
	TIMESTAMP_LTZ TableColumnForIsapPipeColumnDataType
}

func GetTableColumnForIsapPipeColumnDataTypeEnum() TableColumnForIsapPipeColumnDataTypeEnum {
	return TableColumnForIsapPipeColumnDataTypeEnum{
		ROW: TableColumnForIsapPipeColumnDataType{
			value: "ROW",
		},
		MAP_STRING: TableColumnForIsapPipeColumnDataType{
			value: "MAP_STRING",
		},
		MAP_DECIMAL: TableColumnForIsapPipeColumnDataType{
			value: "MAP_DECIMAL",
		},
		TINYINT: TableColumnForIsapPipeColumnDataType{
			value: "TINYINT",
		},
		SMALLINT: TableColumnForIsapPipeColumnDataType{
			value: "SMALLINT",
		},
		INT: TableColumnForIsapPipeColumnDataType{
			value: "INT",
		},
		BIGINT: TableColumnForIsapPipeColumnDataType{
			value: "BIGINT",
		},
		DECIMAL: TableColumnForIsapPipeColumnDataType{
			value: "DECIMAL",
		},
		FLOAT: TableColumnForIsapPipeColumnDataType{
			value: "FLOAT",
		},
		DOUBLE: TableColumnForIsapPipeColumnDataType{
			value: "DOUBLE",
		},
		CHAR: TableColumnForIsapPipeColumnDataType{
			value: "CHAR",
		},
		VARCHAR: TableColumnForIsapPipeColumnDataType{
			value: "VARCHAR",
		},
		STRING: TableColumnForIsapPipeColumnDataType{
			value: "STRING",
		},
		BOOLEAN: TableColumnForIsapPipeColumnDataType{
			value: "BOOLEAN",
		},
		DATE: TableColumnForIsapPipeColumnDataType{
			value: "DATE",
		},
		TIME: TableColumnForIsapPipeColumnDataType{
			value: "TIME",
		},
		TIMESTAMP: TableColumnForIsapPipeColumnDataType{
			value: "TIMESTAMP",
		},
		TIMESTAMP_LTZ: TableColumnForIsapPipeColumnDataType{
			value: "TIMESTAMP_LTZ",
		},
	}
}

func (c TableColumnForIsapPipeColumnDataType) Value() string {
	return c.value
}

func (c TableColumnForIsapPipeColumnDataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableColumnForIsapPipeColumnDataType) UnmarshalJSON(b []byte) error {
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

type TableColumnForIsapPipeColumnType struct {
	value string
}

type TableColumnForIsapPipeColumnTypeEnum struct {
	PHYSICAL         TableColumnForIsapPipeColumnType
	METADATA         TableColumnForIsapPipeColumnType
	VIRTUAL_METADATA TableColumnForIsapPipeColumnType
	COMPUTED         TableColumnForIsapPipeColumnType
}

func GetTableColumnForIsapPipeColumnTypeEnum() TableColumnForIsapPipeColumnTypeEnum {
	return TableColumnForIsapPipeColumnTypeEnum{
		PHYSICAL: TableColumnForIsapPipeColumnType{
			value: "PHYSICAL",
		},
		METADATA: TableColumnForIsapPipeColumnType{
			value: "METADATA",
		},
		VIRTUAL_METADATA: TableColumnForIsapPipeColumnType{
			value: "VIRTUAL_METADATA",
		},
		COMPUTED: TableColumnForIsapPipeColumnType{
			value: "COMPUTED",
		},
	}
}

func (c TableColumnForIsapPipeColumnType) Value() string {
	return c.value
}

func (c TableColumnForIsapPipeColumnType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableColumnForIsapPipeColumnType) UnmarshalJSON(b []byte) error {
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
