package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExportSlowSqlStatisticsRequestBody struct {

	// 数据库类型。慢SQL统计支持的类型：DDS-Community。
	DatastoreType ExportSlowSqlStatisticsRequestBodyDatastoreType `json:"datastore_type"`

	// 开始时间（Unix timestamp），单位：毫秒。
	StartAt int64 `json:"start_at"`

	// 结束时间（Unix timestamp），单位：毫秒。
	EndAt int64 `json:"end_at"`

	// 节点ID列表。
	NodeIds *[]string `json:"node_ids,omitempty"`

	// 统计字段。支持统计的字段：node_id、sql_type、db_name、collection、user、client。默认使用node_id统计。
	StatisticsField *ExportSlowSqlStatisticsRequestBodyStatisticsField `json:"statistics_field,omitempty"`

	// 偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数，默认为20，最大取值100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ExportSlowSqlStatisticsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSlowSqlStatisticsRequestBody struct{}"
	}

	return strings.Join([]string{"ExportSlowSqlStatisticsRequestBody", string(data)}, " ")
}

type ExportSlowSqlStatisticsRequestBodyDatastoreType struct {
	value string
}

type ExportSlowSqlStatisticsRequestBodyDatastoreTypeEnum struct {
	DDS_COMMUNITY ExportSlowSqlStatisticsRequestBodyDatastoreType
}

func GetExportSlowSqlStatisticsRequestBodyDatastoreTypeEnum() ExportSlowSqlStatisticsRequestBodyDatastoreTypeEnum {
	return ExportSlowSqlStatisticsRequestBodyDatastoreTypeEnum{
		DDS_COMMUNITY: ExportSlowSqlStatisticsRequestBodyDatastoreType{
			value: "DDS-Community",
		},
	}
}

func (c ExportSlowSqlStatisticsRequestBodyDatastoreType) Value() string {
	return c.value
}

func (c ExportSlowSqlStatisticsRequestBodyDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportSlowSqlStatisticsRequestBodyDatastoreType) UnmarshalJSON(b []byte) error {
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

type ExportSlowSqlStatisticsRequestBodyStatisticsField struct {
	value string
}

type ExportSlowSqlStatisticsRequestBodyStatisticsFieldEnum struct {
	NODE_ID    ExportSlowSqlStatisticsRequestBodyStatisticsField
	SQL_TYPE   ExportSlowSqlStatisticsRequestBodyStatisticsField
	DB_NAME    ExportSlowSqlStatisticsRequestBodyStatisticsField
	COLLECTION ExportSlowSqlStatisticsRequestBodyStatisticsField
	USER       ExportSlowSqlStatisticsRequestBodyStatisticsField
	CLIENT     ExportSlowSqlStatisticsRequestBodyStatisticsField
}

func GetExportSlowSqlStatisticsRequestBodyStatisticsFieldEnum() ExportSlowSqlStatisticsRequestBodyStatisticsFieldEnum {
	return ExportSlowSqlStatisticsRequestBodyStatisticsFieldEnum{
		NODE_ID: ExportSlowSqlStatisticsRequestBodyStatisticsField{
			value: "node_id",
		},
		SQL_TYPE: ExportSlowSqlStatisticsRequestBodyStatisticsField{
			value: "sql_type",
		},
		DB_NAME: ExportSlowSqlStatisticsRequestBodyStatisticsField{
			value: "db_name",
		},
		COLLECTION: ExportSlowSqlStatisticsRequestBodyStatisticsField{
			value: "collection",
		},
		USER: ExportSlowSqlStatisticsRequestBodyStatisticsField{
			value: "user",
		},
		CLIENT: ExportSlowSqlStatisticsRequestBodyStatisticsField{
			value: "client",
		},
	}
}

func (c ExportSlowSqlStatisticsRequestBodyStatisticsField) Value() string {
	return c.value
}

func (c ExportSlowSqlStatisticsRequestBodyStatisticsField) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportSlowSqlStatisticsRequestBodyStatisticsField) UnmarshalJSON(b []byte) error {
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
