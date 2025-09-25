package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEnhanceFullSqlsRequestBody 查询全量SQL列表请求体
type ListEnhanceFullSqlsRequestBody struct {

	// **参数解释**: 最大查询记录数。例如该参数设定为10，则查询结果最多只显示10条记录。 **约束限制**: 对于公有云25.5.0.1及以上版本，此参数弃用，请勿传值。通过系统系统参数控制最大返回记录数量，默认为200。 **取值范围**: [1, 1000] **默认取值**: 默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 查询开始时间。 **约束限制**: ISO 8601 UTC格式。模式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量。 时区中的+号需要进行URL编码，编码为%2B，时区中的-号无需编码。 例如：北京时间偏移显示为+0800，begin_time=2024-03-15T17:20:33+0800，传参时编码为begin_time=2024-03-15T17:20:33%2B0800。 **取值范围**: 时间区间（begin_time ~ end_time）不能超过30天。 **默认取值**: 不涉及。
	BeginTime string `json:"begin_time"`

	// **参数解释**: 查询结束时间。 **约束限制**: ISO 8601 UTC格式。模式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量。 时区中的+号需要进行URL编码，编码为%2B，时区中的-号无需编码。 例如：北京时间偏移显示为+0800，end_time=2024-03-16T17:20:33+0800，传参时编码为end_time=2024-03-16T17:20:33%2B0800。 **取值范围**: 时间区间（begin_time ~ end_time）不能超过30天。 **默认取值**: 不涉及。
	EndTime string `json:"end_time"`

	// **参数解释**: SQL文本。 **约束限制**: 不涉及。 **取值范围**: 长度1-4096。 **默认取值**: 不涉及。
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

	// **参数解释**: 排序字段。 **约束限制**: 不涉及。 **取值范围**: - begin_time：起始时间。  **默认取值**: begin_time
	OrderBy *ListEnhanceFullSqlsRequestBodyOrderBy `json:"order_by,omitempty"`

	// **参数解释**: 是否为慢SQL。 **约束限制**: 不涉及。 **取值范围**: - true：是慢SQL。 - false：不是慢SQL。  **默认取值**: 不涉及。
	IsSlowSql *bool `json:"is_slow_sql,omitempty"`

	// **参数解释**: 排序方式，支持升序和降序。 **约束限制**: 不涉及。 **取值范围**: - DESC：降序。 - ASC：升序。  **默认取值**: DESC
	Order *ListEnhanceFullSqlsRequestBodyOrder `json:"order,omitempty"`

	// **参数解释**: 字段汇聚查询条件列表。默认取值为[]。 **约束限制**: 不涉及。
	MultiQueries *[]MultiQueryConditionOption `json:"multi_queries,omitempty"`

	// **参数解释**: 组合比较查询条件，可针对某个给定过滤字段，进行区间范围、大小、小于等条件合并查询。默认取值为[]。 **约束限制**: 不涉及。
	CompareConditions *[]CompareConditionOption `json:"compare_conditions,omitempty"`
}

func (o ListEnhanceFullSqlsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnhanceFullSqlsRequestBody struct{}"
	}

	return strings.Join([]string{"ListEnhanceFullSqlsRequestBody", string(data)}, " ")
}

type ListEnhanceFullSqlsRequestBodyOrderBy struct {
	value string
}

type ListEnhanceFullSqlsRequestBodyOrderByEnum struct {
	BEGIN_TIME ListEnhanceFullSqlsRequestBodyOrderBy
}

func GetListEnhanceFullSqlsRequestBodyOrderByEnum() ListEnhanceFullSqlsRequestBodyOrderByEnum {
	return ListEnhanceFullSqlsRequestBodyOrderByEnum{
		BEGIN_TIME: ListEnhanceFullSqlsRequestBodyOrderBy{
			value: "begin_time",
		},
	}
}

func (c ListEnhanceFullSqlsRequestBodyOrderBy) Value() string {
	return c.value
}

func (c ListEnhanceFullSqlsRequestBodyOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEnhanceFullSqlsRequestBodyOrderBy) UnmarshalJSON(b []byte) error {
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

type ListEnhanceFullSqlsRequestBodyOrder struct {
	value string
}

type ListEnhanceFullSqlsRequestBodyOrderEnum struct {
	ASC  ListEnhanceFullSqlsRequestBodyOrder
	DESC ListEnhanceFullSqlsRequestBodyOrder
}

func GetListEnhanceFullSqlsRequestBodyOrderEnum() ListEnhanceFullSqlsRequestBodyOrderEnum {
	return ListEnhanceFullSqlsRequestBodyOrderEnum{
		ASC: ListEnhanceFullSqlsRequestBodyOrder{
			value: "ASC",
		},
		DESC: ListEnhanceFullSqlsRequestBodyOrder{
			value: "DESC",
		},
	}
}

func (c ListEnhanceFullSqlsRequestBodyOrder) Value() string {
	return c.value
}

func (c ListEnhanceFullSqlsRequestBodyOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEnhanceFullSqlsRequestBodyOrder) UnmarshalJSON(b []byte) error {
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
