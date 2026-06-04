package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTablesRequest Request Object
type ListTablesRequest struct {

	// LakeFormation实例ID。创建实例时自动生成。例如：2180518f-42b8-4947-b20b-adfc53981a25。
	InstanceId string `json:"instance_id"`

	// catalog名称。只能包含字母、数字和下划线，且长度为1~256个字符。
	CatalogName string `json:"catalog_name"`

	// 数据库名称。只能包含中文、字母、数字、下划线、中划线，且长度为1~128个字符。
	DatabaseName string `json:"database_name"`

	// 表名称通配符。只能包含中文、字母、数字和_|*.-特殊字符，且长度为1~256个字符。
	TableNamePattern *string `json:"table_name_pattern,omitempty"`

	// 表格式。支持HIVE,ICEBERG,LANCE
	TableFormat *ListTablesRequestTableFormat `json:"table_format,omitempty"`

	// 表类型：MANAGED_TABLE-内表、EXTERNAL_TABLE-外表、VIRTUAL_VIEW-视图、MATERIALIZED_VIEW-物化视图、DICTIONARY_TABLE字典表、LAKE_TABLE内表。
	TableType *ListTablesRequestTableType `json:"table_type,omitempty"`

	// 过滤条件字符串，可以按照属性查询表。 支持的属性查找包括： HIVE_FILTER_FIELD_OWNER HIVE_FILTER_FIELD_LAST_ACCESS HIVE_FILTER_FIELD_PARAMS
	Filter *string `json:"filter,omitempty"`

	// 查询返回条数。默认值为100。最小值为1，最大值为1000。当include_fields只包含name时，最大值可以为5000
	Limit *int32 `json:"limit,omitempty"`

	// 查询的起始记录ID。最小长度为0，最大长度为256。
	Marker *string `json:"marker,omitempty"`

	// 是否倒序查询。
	ReversePage *bool `json:"reverse_page,omitempty"`

	// 是否查询被删除元数据。
	Deleted *bool `json:"deleted,omitempty"`

	// 包含字段，非必填，多个字段使用英文逗号分隔，不填时返回全部字段，当前暂只支持name
	IncludeFields *string `json:"include_fields,omitempty"`
}

func (o ListTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTablesRequest struct{}"
	}

	return strings.Join([]string{"ListTablesRequest", string(data)}, " ")
}

type ListTablesRequestTableFormat struct {
	value string
}

type ListTablesRequestTableFormatEnum struct {
	HIVE    ListTablesRequestTableFormat
	ICEBERG ListTablesRequestTableFormat
	LANCE   ListTablesRequestTableFormat
}

func GetListTablesRequestTableFormatEnum() ListTablesRequestTableFormatEnum {
	return ListTablesRequestTableFormatEnum{
		HIVE: ListTablesRequestTableFormat{
			value: "HIVE",
		},
		ICEBERG: ListTablesRequestTableFormat{
			value: "ICEBERG",
		},
		LANCE: ListTablesRequestTableFormat{
			value: "LANCE",
		},
	}
}

func (c ListTablesRequestTableFormat) Value() string {
	return c.value
}

func (c ListTablesRequestTableFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTablesRequestTableFormat) UnmarshalJSON(b []byte) error {
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

type ListTablesRequestTableType struct {
	value string
}

type ListTablesRequestTableTypeEnum struct {
	MANAGED_TABLE     ListTablesRequestTableType
	EXTERNAL_TABLE    ListTablesRequestTableType
	VIRTUAL_VIEW      ListTablesRequestTableType
	MATERIALIZED_VIEW ListTablesRequestTableType
	DICTIONARY_TABLE  ListTablesRequestTableType
	LAKE_TABLE        ListTablesRequestTableType
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
		DICTIONARY_TABLE: ListTablesRequestTableType{
			value: "DICTIONARY_TABLE",
		},
		LAKE_TABLE: ListTablesRequestTableType{
			value: "LAKE_TABLE",
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
