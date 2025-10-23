package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FullSqlComponetResult 全量SQL组件信息
type FullSqlComponetResult struct {

	// **参数解释**: 组件ID。 **取值范围**: 不涉及。
	ComponentId *string `json:"component_id,omitempty"`

	// **参数解释**: 数据库名称。 **取值范围**: 不涉及。
	DbName *string `json:"db_name,omitempty"`

	// **参数解释**: schema名称。 **取值范围**: 不涉及。
	SchemaName *string `json:"schema_name,omitempty"`

	// **参数解释**: 原始节点。 **取值范围**: 不涉及。
	OriginNode *string `json:"origin_node,omitempty"`

	// **参数解释**: 用户名。 **取值范围**: 不涉及。
	Username *string `json:"username,omitempty"`

	// **参数解释**: 用户发起的请求的应用程序名称。 **取值范围**: 不涉及。
	ApplicationName *string `json:"application_name,omitempty"`

	// **参数解释**: 用户发起的请求的客户端地址。 **取值范围**: 不涉及。
	ClientAddr *string `json:"client_addr,omitempty"`

	// **参数解释**: 用户发起请求的客户端端口。 **取值范围**: 不涉及。
	ClientPort *string `json:"client_port,omitempty"`

	// **参数解释**: 当前语句的外层SQL的归一化SQL ID。 **取值范围**: 不涉及。
	ParentSqlId *string `json:"parent_sql_id,omitempty"`

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

	// **参数解释**: 线程ID。 **取值范围**: 不涉及。
	ThreadId *string `json:"thread_id,omitempty"`

	// **参数解释**: 会话ID。 **取值范围**: 不涉及。
	SessionId *string `json:"session_id,omitempty"`

	// **参数解释**: 开始时间，格式为“yyyy-mm-ddThh:mm:ss.SSSSSZ”。 **取值范围**: 不涉及。
	BeginTime *string `json:"begin_time,omitempty"`

	// **参数解释**: 结束时间，格式为“yyyy-mm-ddThh:mm:ss.SSSSSZ”。 **取值范围**: 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**: 慢SQL阈值。 **取值范围**: 不涉及。
	SlowQueryThreshold *int64 `json:"slow_query_threshold,omitempty"`

	// **参数解释**: 软解析次数。 **取值范围**: 不涉及。
	NSoftParse *int64 `json:"n_soft_parse,omitempty"`

	// **参数解释**: 硬解析次数。 **取值范围**: 不涉及。
	NHardParse *int64 `json:"n_hard_parse,omitempty"`

	// **参数解释**: 执行计划。 **取值范围**: 不涉及。
	QueryPlan *string `json:"query_plan,omitempty"`

	// **参数解释**: SELECT语句的返回结果集行数。 **取值范围**: 不涉及。
	NReturnRows *int64 `json:"n_return_rows,omitempty"`

	// **参数解释**: 随机扫描行。 **取值范围**: 不涉及。
	NTuplesFetched *int64 `json:"n_tuples_fetched,omitempty"`

	// **参数解释**: 顺序扫描行。 **取值范围**: 不涉及。
	NTuplesReturned *int64 `json:"n_tuples_returned,omitempty"`

	// **参数解释**: 插入行。 **取值范围**: 不涉及。
	NTuplesInserted *int64 `json:"n_tuples_inserted,omitempty"`

	// **参数解释**: 更新行。 **取值范围**: 不涉及。
	NTuplesUpdated *int64 `json:"n_tuples_updated,omitempty"`

	// **参数解释**: 删除行。 **取值范围**: 不涉及。
	NTuplesDeleted *int64 `json:"n_tuples_deleted,omitempty"`

	// **参数解释**: buffer的块访问次数。 **取值范围**: 不涉及。\"\"
	NBlocksFetched *int64 `json:"n_blocks_fetched,omitempty"`

	// **参数解释**: buffer的块命中次数。 **取值范围**: 不涉及。
	NBlocksHit *int64 `json:"n_blocks_hit,omitempty"`

	// **参数解释**: 有效DB时间花费。 **取值范围**: 不涉及。
	DbTime *int64 `json:"db_time,omitempty"`

	// **参数解释**: CPU时间（单位：微秒）。 **取值范围**: 不涉及。
	CpuTime *int64 `json:"cpu_time,omitempty"`

	// **参数解释**: 执行器内执行时间（单位：微秒）。 **取值范围**: 不涉及。
	ExecutionTime *int64 `json:"execution_time,omitempty"`

	// **参数解释**: SQL解析时间（单位：微秒）。 **取值范围**: 不涉及。
	ParseTime *int64 `json:"parse_time,omitempty"`

	// **参数解释**: 执行器内执行时间（单位：微秒）。 **取值范围**: 不涉及。
	PlanTime *int64 `json:"plan_time,omitempty"`

	// **参数解释**: SQL重写时间（单位：微秒）。 **取值范围**: 不涉及。
	RewriteTime *int64 `json:"rewrite_time,omitempty"`

	// **参数解释**: plpgsql上的执行时间（单位：微秒）。 **取值范围**: 不涉及。
	PlExecutionTime *int64 `json:"pl_execution_time,omitempty"`

	// **参数解释**: plpgsql上的编译时间（单位：微秒）。 **取值范围**: 不涉及。
	PlCompilationTime *int64 `json:"pl_compilation_time,omitempty"`

	// **参数解释**: IO时间（单位：微秒）。 **取值范围**: 不涉及。
	DataIoTime *int64 `json:"data_io_time,omitempty"`

	// **参数解释**: 加锁次数。 **取值范围**: 不涉及。
	LockCount *int64 `json:"lock_count,omitempty"`

	// **参数解释**: 加锁耗时。 **取值范围**: 不涉及。
	LockTime *int64 `json:"lock_time,omitempty"`

	// **参数解释**: 加锁等待次数。 **取值范围**: 不涉及。
	LockWaitCount *int64 `json:"lock_wait_count,omitempty"`

	// **参数解释**: 加锁等待时间。 **取值范围**: 不涉及。
	LockWaitTime *int64 `json:"lock_wait_time,omitempty"`

	// **参数解释**: 详细列表。 **取值范围**: 不涉及。
	Details *string `json:"details,omitempty"`

	// **参数解释**: 是否慢SQL。 **取值范围**: 不涉及。
	IsSlowSql *bool `json:"is_slow_sql,omitempty"`

	// **参数解释**: 可能导致该SQL为慢SQL的风险信息。 **取值范围**: 不涉及。
	Advise *string `json:"advise,omitempty"`

	// **参数解释**: 语句完成状态。 **取值范围**: 不涉及。
	FinishStatus *string `json:"finish_status,omitempty"`

	// **参数解释**: 通过物理连接发送消息的网络状态。 **取值范围**: 不涉及。
	NetSendInfo *string `json:"net_send_info,omitempty"`

	// **参数解释**: 通过物理连接接收消息的网络状态。 **取值范围**: 不涉及。
	NetRecvInfo *string `json:"net_recv_info,omitempty"`

	// **参数解释**: 通过逻辑连接发送消息的网络状态。 **取值范围**: 不涉及。
	NetStreamSendInfo *string `json:"net_stream_send_info,omitempty"`

	// **参数解释**: 通过逻辑连接接收消息的网络状态。 **取值范围**: 不涉及。
	NetStreamRecvInfo *string `json:"net_stream_recv_info,omitempty"`
}

func (o FullSqlComponetResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullSqlComponetResult struct{}"
	}

	return strings.Join([]string{"FullSqlComponetResult", string(data)}, " ")
}
