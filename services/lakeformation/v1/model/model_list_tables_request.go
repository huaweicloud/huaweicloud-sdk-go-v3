package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListTablesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名字通配符
	TableNamePattern *string `json:"table_name_pattern,omitempty"`

	// table_type
	TableType *ListTablesRequestTableType `json:"table_type,omitempty"`

	// 过滤条件字符串，可以按照属性查询表。 支持的属性查找包括： HIVE_FILTER_FIELD_OWNER HIVE_FILTER_FIELD_LAST_ACCESS HIVE_FILTER_FIELD_PARAMS
	Filter *string `json:"filter,omitempty"`

	// 返回的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 查询的起始记录ID
	Marker *string `json:"marker,omitempty"`

	// 是否查询上一页
	ReversePage *bool `json:"reverse_page,omitempty"`
}

func (o ListTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTablesRequest struct{}"
	}

	return strings.Join([]string{"ListTablesRequest", string(data)}, " ")
}

type ListTablesRequestTableType struct {
	value string
}

type ListTablesRequestTableTypeEnum struct {
	MANAGED_TABLE     ListTablesRequestTableType
	EXTERNAL_TABLE    ListTablesRequestTableType
	VIRTUAL_VIEW      ListTablesRequestTableType
	MATERIALIZED_VIEW ListTablesRequestTableType
}

func GetListTablesRequestTableTypeEnum() ListTablesRequestTableTypeEnum {
	return ListTablesRequestTableTypeEnum{
		MANAGED_TABLE: ListTablesRequestTableType{
			value: "MANAGED_TABLE",
		},
		EXTERNAL_TABLE: ListTablesRequestTableType{
			value: "EXTERNAL_TABLE",
		},
		VIRTUAL_VIEW: ListTablesRequestTableType{
			value: "VIRTUAL_VIEW",
		},
		MATERIALIZED_VIEW: ListTablesRequestTableType{
			value: "MATERIALIZED_VIEW",
		},
	}
}

func (c ListTablesRequestTableType) Value() string {
	return c.value
}

func (c ListTablesRequestTableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTablesRequestTableType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
