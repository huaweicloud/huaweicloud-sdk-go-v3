package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TableMeta 表描述信息
type TableMeta struct {

	// 表名
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名字
	TableName string `json:"table_name"`

	// 表类型
	TableType TableMetaTableType `json:"table_type"`

	// 表描述信息
	Comments string `json:"comments"`

	// 分区列以外的所有字段。
	Columns *[]Column `json:"columns,omitempty"`

	// 分区列的信息。
	PartitionKeys *[]Column `json:"partition_keys,omitempty"`
}

func (o TableMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableMeta struct{}"
	}

	return strings.Join([]string{"TableMeta", string(data)}, " ")
}

type TableMetaTableType struct {
	value string
}

type TableMetaTableTypeEnum struct {
	MANAGED_TABLE     TableMetaTableType
	EXTERNAL_TABLE    TableMetaTableType
	VIRTUAL_VIEW      TableMetaTableType
	MATERIALIZED_VIEW TableMetaTableType
}

func GetTableMetaTableTypeEnum() TableMetaTableTypeEnum {
	return TableMetaTableTypeEnum{
		MANAGED_TABLE: TableMetaTableType{
			value: "MANAGED_TABLE",
		},
		EXTERNAL_TABLE: TableMetaTableType{
			value: "EXTERNAL_TABLE",
		},
		VIRTUAL_VIEW: TableMetaTableType{
			value: "VIRTUAL_VIEW",
		},
		MATERIALIZED_VIEW: TableMetaTableType{
			value: "MATERIALIZED_VIEW",
		},
	}
}

func (c TableMetaTableType) Value() string {
	return c.value
}

func (c TableMetaTableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableMetaTableType) UnmarshalJSON(b []byte) error {
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
