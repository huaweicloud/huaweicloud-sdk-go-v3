package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListEnhanceFullSqlStatisticsRequestBody struct {

	// **参数解释**: 查询记录数。例如该参数设定为10，则查询结果最多只显示10条记录。 **约束限制**: 不涉及。 **取值范围**: [1, 1000] **默认取值**: 默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。例如：该参数指定为0，limit指定为10，则只展示第1~10条数据。 **约束限制**: 不涉及。 **取值范围**: [0, 9223372036854774807] **默认取值**: 默认为0（偏移0条数据，表示从第一条数据开始查询）。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 节点ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	NodeId *string `json:"node_id,omitempty"`

	// **参数解释**: 查询开始时间。 **约束限制**: ISO 8601 UTC格式。模式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量。时区中的+号需要进行URL编码，编码为%2B，时区中的-号无需编码。例如：北京时间偏移显示为+0800，begin_time=2024-03-15T17:20:33+0800，传参时编码为begin_time=2024-03-15T17:20:33%2B0800。 **取值范围**: 时间区间（begin_time ~ end_time）不能超过30天。 **默认取值**: 不涉及。
	BeginTime string `json:"begin_time"`

	// **参数解释**: 查询结束时间。 **约束限制**: ISO 8601 UTC格式。模式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量。时区中的+号需要进行URL编码，编码为%2B，时区中的-号无需编码。例如：北京时间偏移显示为+0800，end_time=2024-03-16T17:20:33+0800，传参时编码为end_time=2024-03-16T17:20:33%2B0800。 **取值范围**: 时间区间（begin_time ~ end_time）不能超过30天。 **默认取值**: 不涉及。
	EndTime string `json:"end_time"`

	// **参数解释**: SQL文本。 **约束限制**: 不涉及。 **取值范围**: 长度1~4096。 **默认取值**: 不涉及。
	Query *string `json:"query,omitempty"`

	// **参数解释**: 归一化SQL ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SqlId *string `json:"sql_id,omitempty"`

	// **参数解释**: 唯一SQL ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SqlExecId *string `json:"sql_exec_id,omitempty"`

	// **参数解释**: 事务ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	TransactionId *string `json:"transaction_id,omitempty"`

	// **参数解释**: 链路ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	TraceId *string `json:"trace_id,omitempty"`

	// **参数解释**: 数据库名称。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	DbName *string `json:"db_name,omitempty"`

	// **参数解释**: schema名称。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SchemaName *string `json:"schema_name,omitempty"`

	// **参数解释**: 用户名称。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	Username *string `json:"username,omitempty"`

	// **参数解释**: 客户端地址。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	ClientAddr *string `json:"client_addr,omitempty"`

	// **参数解释**: 客户端端口。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	ClientPort *string `json:"client_port,omitempty"`

	// **参数解释**: 排序字段。 **约束限制**: 不涉及。 **取值范围**: - sql_id - sql_count - avg_db_time - avg_cpu_time - avg_execution_time - avg_data_io_time - start_time_stamp  **默认取值**: sql_count
	OrderBy *ListEnhanceFullSqlStatisticsRequestBodyOrderBy `json:"order_by,omitempty"`

	// **参数解释**: 是否是慢SQL。 **约束限制**: 不涉及。 **取值范围**: - true：是慢SQL。 - false：不是慢SQL。  **默认取值**: 不涉及。
	IsSlowSql *bool `json:"is_slow_sql,omitempty"`

	// **参数解释**: 排序方式，支持升序和降序。 **约束限制**: 不涉及。 **取值范围**: - DESC：降序。 - ASC：升序。  **默认取值**: DESC
	Order *ListEnhanceFullSqlStatisticsRequestBodyOrder `json:"order,omitempty"`

	// **参数解释**: 应用名称。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	ApplicationName *string `json:"application_name,omitempty"`

	// **参数解释**: 组件ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	ComponentId *string `json:"component_id,omitempty"`

	// **参数解释**: 字段汇聚查询条件列表。 **约束限制**: 只支持针对query字段全与或者全或的查询。
	MultiQueries *[]MultiQueryConditionOption `json:"multi_queries,omitempty"`

	// **参数解释**: 组合比较查询条件，可针对某个给定过滤字段，进行区间范围、大小、小于等条件合并查询。 **约束限制**: 不涉及。
	CompareConditions *[]CompareConditionOption `json:"compare_conditions,omitempty"`
}

func (o ListEnhanceFullSqlStatisticsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnhanceFullSqlStatisticsRequestBody struct{}"
	}

	return strings.Join([]string{"ListEnhanceFullSqlStatisticsRequestBody", string(data)}, " ")
}

type ListEnhanceFullSqlStatisticsRequestBodyOrderBy struct {
	value string
}

type ListEnhanceFullSqlStatisticsRequestBodyOrderByEnum struct {
	SQL_ID             ListEnhanceFullSqlStatisticsRequestBodyOrderBy
	SQL_COUNT          ListEnhanceFullSqlStatisticsRequestBodyOrderBy
	AVG_DB_TIME        ListEnhanceFullSqlStatisticsRequestBodyOrderBy
	AVG_CPU_TIME       ListEnhanceFullSqlStatisticsRequestBodyOrderBy
	AVG_EXECUTION_TIME ListEnhanceFullSqlStatisticsRequestBodyOrderBy
	AVG_DATA_IO_TIME   ListEnhanceFullSqlStatisticsRequestBodyOrderBy
	START_TIME_STAMP   ListEnhanceFullSqlStatisticsRequestBodyOrderBy
}

func GetListEnhanceFullSqlStatisticsRequestBodyOrderByEnum() ListEnhanceFullSqlStatisticsRequestBodyOrderByEnum {
	return ListEnhanceFullSqlStatisticsRequestBodyOrderByEnum{
		SQL_ID: ListEnhanceFullSqlStatisticsRequestBodyOrderBy{
			value: "sql_id",
		},
		SQL_COUNT: ListEnhanceFullSqlStatisticsRequestBodyOrderBy{
			value: "sql_count",
		},
		AVG_DB_TIME: ListEnhanceFullSqlStatisticsRequestBodyOrderBy{
			value: "avg_db_time",
		},
		AVG_CPU_TIME: ListEnhanceFullSqlStatisticsRequestBodyOrderBy{
			value: "avg_cpu_time",
		},
		AVG_EXECUTION_TIME: ListEnhanceFullSqlStatisticsRequestBodyOrderBy{
			value: "avg_execution_time",
		},
		AVG_DATA_IO_TIME: ListEnhanceFullSqlStatisticsRequestBodyOrderBy{
			value: "avg_data_io_time",
		},
		START_TIME_STAMP: ListEnhanceFullSqlStatisticsRequestBodyOrderBy{
			value: "start_time_stamp",
		},
	}
}

func (c ListEnhanceFullSqlStatisticsRequestBodyOrderBy) Value() string {
	return c.value
}

func (c ListEnhanceFullSqlStatisticsRequestBodyOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEnhanceFullSqlStatisticsRequestBodyOrderBy) UnmarshalJSON(b []byte) error {
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

type ListEnhanceFullSqlStatisticsRequestBodyOrder struct {
	value string
}

type ListEnhanceFullSqlStatisticsRequestBodyOrderEnum struct {
	DESC ListEnhanceFullSqlStatisticsRequestBodyOrder
	ASC  ListEnhanceFullSqlStatisticsRequestBodyOrder
}

func GetListEnhanceFullSqlStatisticsRequestBodyOrderEnum() ListEnhanceFullSqlStatisticsRequestBodyOrderEnum {
	return ListEnhanceFullSqlStatisticsRequestBodyOrderEnum{
		DESC: ListEnhanceFullSqlStatisticsRequestBodyOrder{
			value: "DESC",
		},
		ASC: ListEnhanceFullSqlStatisticsRequestBodyOrder{
			value: "ASC",
		},
	}
}

func (c ListEnhanceFullSqlStatisticsRequestBodyOrder) Value() string {
	return c.value
}

func (c ListEnhanceFullSqlStatisticsRequestBodyOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEnhanceFullSqlStatisticsRequestBodyOrder) UnmarshalJSON(b []byte) error {
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
