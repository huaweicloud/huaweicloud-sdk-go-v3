package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RealTimeSessionResult struct {

	// **参数解释**： 会话ID。 **取值范围**： 不涉及。
	SessionId *string `json:"session_id,omitempty"`

	// **参数解释**： 线程ID。 **取值范围**： 不涉及。
	Pid *string `json:"pid,omitempty"`

	// **参数解释**： SQL ID。 **取值范围**： 不涉及。
	UniqueSqlId *string `json:"unique_sql_id,omitempty"`

	// **参数解释**： 数据库名称。 **取值范围**： 不涉及。
	DatabaseName *string `json:"database_name,omitempty"`

	// **参数解释**： 客户端IP。 **取值范围**： 不涉及。
	ClientIp *string `json:"client_ip,omitempty"`

	// **参数解释**： 用户名称。 **取值范围**： 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**： 是否等待。 **取值范围**： 不涉及。
	Wait *string `json:"wait,omitempty"`

	// **参数解释**： 阻塞会话。 **取值范围**： 不涉及。
	BlockSession *string `json:"block_session,omitempty"`

	// **参数解释**： 等待事件。 **取值范围**： 不涉及。
	WaitEvent *string `json:"wait_event,omitempty"`

	// **参数解释**： 状态。 **取值范围**： 不涉及。
	State *string `json:"state,omitempty"`

	// **参数解释**： 语句执行时长。 **取值范围**： 不涉及。
	QueryRuntime *string `json:"query_runtime,omitempty"`

	// **参数解释**： SQL文本。 **取值范围**： 不涉及。
	Query *string `json:"query,omitempty"`

	// **参数解释**： 会话开始时间。 **取值范围**： 不涉及。
	BackEndStart *int32 `json:"back_end_start,omitempty"`

	// **参数解释**： 事务开始时间。 **取值范围**： 不涉及。
	TransactionStart *int32 `json:"transaction_start,omitempty"`

	// **参数解释**： 语句开始时间。 **取值范围**： 不涉及。
	QueryStart *int32 `json:"query_start,omitempty"`

	// **参数解释**： 应用名称。 **取值范围**： 不涉及。
	ApplicationName *string `json:"application_name,omitempty"`

	// **参数解释**： 会话执行时长（单位：秒）。 **取值范围**： 不涉及。
	ExecTime *string `json:"exec_time,omitempty"`

	// **参数解释**： 会话建立事务的数量。 **取值范围**： 不涉及。
	TransNum *string `json:"trans_num,omitempty"`

	// **参数解释**： 会话中事务回滚的次数。 **取值范围**： 不涉及。
	RollbackNum *string `json:"rollback_num,omitempty"`

	// **参数解释**： 会话执行的sql数。 **取值范围**： 不涉及。
	SqlNum *string `json:"sql_num,omitempty"`

	// **参数解释**： 客户端用于与后台通讯的TCP端口号，如果使用Unix套接字，则为-1。 **取值范围**： 不涉及。
	ClientPort *string `json:"client_port,omitempty"`

	// **参数解释**： 会话执行的sql数。 **取值范围**： 不涉及。
	QueryId *string `json:"query_id,omitempty"`

	// **参数解释**： 当前用户上一次执行的事务持续时间。 **取值范围**： 不涉及。
	TransactionTimeCost *string `json:"transaction_time_cost,omitempty"`

	// **参数解释**： 驱动传入的trace id，用于标识应用的一次请求。 **取值范围**： 不涉及。
	TraceId *string `json:"trace_id,omitempty"`

	// **参数解释**： 当前用户上次执行的全局会话ID。 **取值范围**： 不涉及。
	GlobalSessionId *string `json:"global_session_id,omitempty"`

	// **参数解释**： 当前用户上次执行的顶层事务ID。 **取值范围**： 不涉及。
	TopTransactionId *string `json:"top_transaction_id,omitempty"`

	// **参数解释**： 当前用户上次执行的事务ID。 **取值范围**： 不涉及。
	CurrentTransactionId *string `json:"current_transaction_id,omitempty"`

	// **参数解释**： 当前用户上次执行的事务使用的XLOG量，易读格式。 **取值范围**： 不涉及。
	XlogQuantityPretty *string `json:"xlog_quantity_pretty,omitempty"`

	// **参数解释**： 实例线程等待状态。 **取值范围**： 不涉及。
	WaitStatus *string `json:"wait_status,omitempty"`

	// **参数解释**： 实例线程的轻量级线程号。 **取值范围**： 不涉及。
	LwtId *string `json:"lwt_id,omitempty"`

	// **参数解释**： 实例线程名。 **取值范围**： 不涉及。
	ThreadName *string `json:"thread_name,omitempty"`

	// **参数解释**： 实例等锁模式。 **取值范围**： 不涉及。
	LockMode *string `json:"lock_mode,omitempty"`

	// **参数解释**： 实例父会话ID。 **取值范围**： 不涉及。
	ParentSessionId *string `json:"parent_session_id,omitempty"`

	// **参数解释**： 实例并行线程的ID。 **取值范围**： 不涉及。
	SmpId *string `json:"smp_id,omitempty"`

	// **参数解释**： 实例线程正等待获取的锁的信息。 **取值范围**： 不涉及。
	LockTag *string `json:"lock_tag,omitempty"`

	// **参数解释**： 组件名称。 **取值范围**： 不涉及。
	ComponentName *string `json:"component_name,omitempty"`
}

func (o RealTimeSessionResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RealTimeSessionResult struct{}"
	}

	return strings.Join([]string{"RealTimeSessionResult", string(data)}, " ")
}
