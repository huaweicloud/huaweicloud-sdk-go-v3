package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FullSqlResult **参数解释**: 全量SQL。
type FullSqlResult struct {

	// **参数解释**: SQL记录唯一键ID。 **取值范围**: 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**: 实例ID。 **取值范围**: 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**: 节点ID。 **取值范围**: 不涉及。
	NodeId *string `json:"node_id,omitempty"`

	// **参数解释**: 组件ID。 **取值范围**: 不涉及。
	ComponentId *string `json:"component_id,omitempty"`

	// **参数解释**: 数据库名称。 **取值范围**: 不涉及。
	DbName *string `json:"db_name,omitempty"`

	// **参数解释**: schema名称。 **取值范围**: 不涉及。
	SchemaName *string `json:"schema_name,omitempty"`

	// **参数解释**: 用户名称。 **取值范围**: 不涉及。
	Username *string `json:"username,omitempty"`

	// **参数解释**: 用户发起的请求的应用程序名称。 **取值范围**: 不涉及。
	ApplicationName *string `json:"application_name,omitempty"`

	// **参数解释**: 用户发起的请求的客户端地址。 **取值范围**: 不涉及。
	ClientAddr *string `json:"client_addr,omitempty"`

	// **参数解释**: 用户发起请求的客户端端口。 **取值范围**: 不涉及。
	ClientPort *string `json:"client_port,omitempty"`

	// **参数解释**: 归一化SQL ID，对应内核字段：unique_sql_id。 **取值范围**: 不涉及。
	SqlId *string `json:"sql_id,omitempty"`

	// **参数解释**: 唯一SQL ID，对应内核字段：debug_query_id。 **取值范围**: 不涉及。
	SqlExecId *string `json:"sql_exec_id,omitempty"`

	// **参数解释**: 事务ID，对应内核字段：debug_query_id。 **取值范围**: 不涉及。
	TransactionId *string `json:"transaction_id,omitempty"`

	// **参数解释**: 链路ID。 **取值范围**: 不涉及。
	TraceId *string `json:"trace_id,omitempty"`

	// **参数解释**: 归一化SQL。 **取值范围**: 不涉及。
	Query *string `json:"query,omitempty"`

	// **参数解释**: 解析后的原始SQL文本。 **取值范围**: 开启track_stmt_parameter参数后，会把SQL文本中的变量替换成真实值，展示原始的SQL。对于track_stmt_parameter参数关闭时采集的SQL文本，无法获取到SQL参数变量的值，展示的内容为空。
	Sql *string `json:"sql,omitempty"`

	// **参数解释**: 开始时间，格式为“yyyy-mm-ddThh:mm:ss.SSSSSZ”。 **取值范围**: 不涉及。
	BeginTime *string `json:"begin_time,omitempty"`

	// **参数解释**: 结束时间，格式为“yyyy-mm-ddThh:mm:ss.SSSSSZ”。 **取值范围**: 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**: 执行总时间（单位：微秒）。 **取值范围**: 不涉及。
	AllTime *int64 `json:"all_time,omitempty"`

	// **参数解释**: 有效DB时间（单位：微秒）。 **取值范围**: 不涉及。
	DbTime *int64 `json:"db_time,omitempty"`

	// **参数解释**: CPU时间（单位：微秒）。 **取值范围**: 不涉及。
	CpuTime *int64 `json:"cpu_time,omitempty"`

	// **参数解释**: IO时间（单位：微秒）。 **取值范围**: 不涉及。
	DataIoTime *int64 `json:"data_io_time,omitempty"`

	// **参数解释**: 执行器内执行时间（单位：微秒）。 **取值范围**: 不涉及。
	ExecutionTime *int64 `json:"execution_time,omitempty"`

	// **参数解释**: 扫描行。 **取值范围**: 不涉及。
	ScanLines *int64 `json:"scan_lines,omitempty"`

	// **参数解释**: 插入行。 **取值范围**: 不涉及。
	InsertRows *int64 `json:"insert_rows,omitempty"`

	// **参数解释**: 更新行。 **取值范围**: 不涉及。
	UpdateRows *int64 `json:"update_rows,omitempty"`

	// **参数解释**: 删除行。 **取值范围**: 不涉及。
	DeleteRows *int64 `json:"delete_rows,omitempty"`

	// **参数解释**: 是否慢SQL。 **取值范围**: 不涉及。
	IsSlowSql *bool `json:"is_slow_sql,omitempty"`

	// **参数解释**: SQL开始时间。格式为13位标准时间戳，如1754647168354。 **取值范围**: 不涉及。
	StartTimestamp *int64 `json:"start_timestamp,omitempty"`

	// **参数解释**: SQL结束时间，格式为13位标准时间戳，如1754647168355。 **取值范围**: 不涉及。
	FinishTimestamp *int64 `json:"finish_timestamp,omitempty"`

	// **参数解释**: SQL命中率。 计划即将下线，请勿使用。 **取值范围**: 不涉及。
	HitRate *float64 `json:"hit_rate,omitempty"`
}

func (o FullSqlResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullSqlResult struct{}"
	}

	return strings.Join([]string{"FullSqlResult", string(data)}, " ")
}
