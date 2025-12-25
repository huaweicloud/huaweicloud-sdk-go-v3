package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type IsapTableColumn struct {

	// 表格列名称
	ColumnName *string `json:"column_name,omitempty"`

	// **参数解释**: 列字段类型 - PHYSICAL 物理字段 - METADATA 元数据字段 - VIRTUAL_METADATA 虚拟元数据字段 - COMPUTED 计算字段  **约束限制** 不涉及  **取值范围**: - PHYSICAL - METADATA - VIRTUAL_METADATA - COMPUTED  **默认值** 不涉及
	ColumnType *IsapTableColumnColumnType `json:"column_type,omitempty"`

	// 表格列类型设置
	ColumnTypeSetting *string `json:"column_type_setting,omitempty"`

	// **参数解释**: 列字段数据类型 - ROW 行类型 - MAP_STRING 字符串映射类型 - MAP_DECIMAL 小数映射类型 - TINYINT 微整型 - SMALLINT 小整型 - INT 整型 - BIGINT 长整型 - DECIMAL 精确小数类型 - FLOAT 单精度浮点数 - DOUBLE 双精度浮点数 - CHAR 定长字符串 - VARCHAR 不定长字符串 - STRING 字符串类型 - KEYWORD 关键字类型 - BOOLEAN 布尔类型 - DATE 日期类型 - TIME 时间类型 - TIMESTAMP 时间戳类型 - TIMESTAMP_LTZ 本地时间戳类型  **约束限制** 不涉及  **取值范围**: - ROW - MAP_STRING - MAP_DECIMAL - TINYINT - SMALLINT - INT - BIGINT - DECIMAL - FLOAT - DOUBLE - CHAR - VARCHAR - STRING - KEYWORD - BOOLEAN - DATE - TIME - TIMESTAMP - TIMESTAMP_LTZ  **默认值** 不涉及
	ColumnDataType *IsapTableColumnColumnDataType `json:"column_data_type,omitempty"`

	// 表格列数据类型设置
	ColumnDataTypeSetting *string `json:"column_data_type_setting,omitempty"`

	// 是否为空
	Nullable *bool `json:"nullable,omitempty"`

	// 是否为数组
	Array *bool `json:"array,omitempty"`

	// 深度
	Depth *int32 `json:"depth,omitempty"`

	// 父级名称
	ParentName *string `json:"parent_name,omitempty"`

	// 所属名称
	OwnName *string `json:"own_name,omitempty"`

	ColumnDisplaySetting *ColumnDisplaySetting `json:"column_display_setting,omitempty"`

	// 列序号
	ColumnSequenceNumber *int32 `json:"column_sequence_number,omitempty"`
}

func (o IsapTableColumn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapTableColumn struct{}"
	}

	return strings.Join([]string{"IsapTableColumn", string(data)}, " ")
}

type IsapTableColumnColumnType struct {
	value string
}

type IsapTableColumnColumnTypeEnum struct {
	PHYSICAL         IsapTableColumnColumnType
	METADATA         IsapTableColumnColumnType
	VIRTUAL_METADATA IsapTableColumnColumnType
	COMPUTED         IsapTableColumnColumnType
}

func GetIsapTableColumnColumnTypeEnum() IsapTableColumnColumnTypeEnum {
	return IsapTableColumnColumnTypeEnum{
		PHYSICAL: IsapTableColumnColumnType{
			value: "PHYSICAL",
		},
		METADATA: IsapTableColumnColumnType{
			value: "METADATA",
		},
		VIRTUAL_METADATA: IsapTableColumnColumnType{
			value: "VIRTUAL_METADATA",
		},
		COMPUTED: IsapTableColumnColumnType{
			value: "COMPUTED",
		},
	}
}

func (c IsapTableColumnColumnType) Value() string {
	return c.value
}

func (c IsapTableColumnColumnType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IsapTableColumnColumnType) UnmarshalJSON(b []byte) error {
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

type IsapTableColumnColumnDataType struct {
	value string
}

type IsapTableColumnColumnDataTypeEnum struct {
	ROW           IsapTableColumnColumnDataType
	MAP_STRING    IsapTableColumnColumnDataType
	MAP_DECIMAL   IsapTableColumnColumnDataType
	TINYINT       IsapTableColumnColumnDataType
	SMALLINT      IsapTableColumnColumnDataType
	INT           IsapTableColumnColumnDataType
	BIGINT        IsapTableColumnColumnDataType
	DECIMAL       IsapTableColumnColumnDataType
	FLOAT         IsapTableColumnColumnDataType
	DOUBLE        IsapTableColumnColumnDataType
	CHAR          IsapTableColumnColumnDataType
	VARCHAR       IsapTableColumnColumnDataType
	STRING        IsapTableColumnColumnDataType
	KEYWORD       IsapTableColumnColumnDataType
	BOOLEAN       IsapTableColumnColumnDataType
	DATE          IsapTableColumnColumnDataType
	TIME          IsapTableColumnColumnDataType
	TIMESTAMP     IsapTableColumnColumnDataType
	TIMESTAMP_LTZ IsapTableColumnColumnDataType
}

func GetIsapTableColumnColumnDataTypeEnum() IsapTableColumnColumnDataTypeEnum {
	return IsapTableColumnColumnDataTypeEnum{
		ROW: IsapTableColumnColumnDataType{
			value: "ROW",
		},
		MAP_STRING: IsapTableColumnColumnDataType{
			value: "MAP_STRING",
		},
		MAP_DECIMAL: IsapTableColumnColumnDataType{
			value: "MAP_DECIMAL",
		},
		TINYINT: IsapTableColumnColumnDataType{
			value: "TINYINT",
		},
		SMALLINT: IsapTableColumnColumnDataType{
			value: "SMALLINT",
		},
		INT: IsapTableColumnColumnDataType{
			value: "INT",
		},
		BIGINT: IsapTableColumnColumnDataType{
			value: "BIGINT",
		},
		DECIMAL: IsapTableColumnColumnDataType{
			value: "DECIMAL",
		},
		FLOAT: IsapTableColumnColumnDataType{
			value: "FLOAT",
		},
		DOUBLE: IsapTableColumnColumnDataType{
			value: "DOUBLE",
		},
		CHAR: IsapTableColumnColumnDataType{
			value: "CHAR",
		},
		VARCHAR: IsapTableColumnColumnDataType{
			value: "VARCHAR",
		},
		STRING: IsapTableColumnColumnDataType{
			value: "STRING",
		},
		KEYWORD: IsapTableColumnColumnDataType{
			value: "KEYWORD",
		},
		BOOLEAN: IsapTableColumnColumnDataType{
			value: "BOOLEAN",
		},
		DATE: IsapTableColumnColumnDataType{
			value: "DATE",
		},
		TIME: IsapTableColumnColumnDataType{
			value: "TIME",
		},
		TIMESTAMP: IsapTableColumnColumnDataType{
			value: "TIMESTAMP",
		},
		TIMESTAMP_LTZ: IsapTableColumnColumnDataType{
			value: "TIMESTAMP_LTZ",
		},
	}
}

func (c IsapTableColumnColumnDataType) Value() string {
	return c.value
}

func (c IsapTableColumnColumnDataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IsapTableColumnColumnDataType) UnmarshalJSON(b []byte) error {
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
