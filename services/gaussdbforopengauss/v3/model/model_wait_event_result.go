package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WaitEventResult struct {

	// **参数解释**: 会话ID。 **取值范围**: 不涉及。
	SessionId *string `json:"session_id,omitempty"`

	// **参数解释**: 线程ID。 **取值范围**: 不涉及。
	Pid *string `json:"pid,omitempty"`

	// **参数解释**: 等待事件。 参见“开发指南 > 系统表和系统视图 > 系统视图 > 其他系统视图 > PG_THREAD_WAIT_STATUS”中的wait_event字段。 **取值范围**: 不涉及。
	WaitEvent *string `json:"wait_event,omitempty"`

	// **参数解释**: 等待状态。 参见“开发指南 > 系统表和系统视图 > 系统视图 > 其他系统视图 > PG_THREAD_WAIT_STATUS”中的等待状态列表。 **取值范围**: 不涉及。
	WaitStatus *string `json:"wait_status,omitempty"`

	// **参数解释**: 数据库名称。 **取值范围**: 不涉及。
	DatabaseName *string `json:"database_name,omitempty"`

	// **参数解释**: 用户名称。 **取值范围**: 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: 客户端地址。 **取值范围**: 不涉及。
	ClientAddr *string `json:"client_addr,omitempty"`

	// **参数解释**: 客户端用于与后台通讯的TCP端口号。 **取值范围**: 不涉及。
	ClientPort *string `json:"client_port,omitempty"`

	// **参数解释**: 会话持续时间，单位：秒。 **取值范围**: 不涉及。
	SessionTime *string `json:"session_time,omitempty"`

	// **参数解释**: 会话开始时间。 **取值范围**: 不涉及。
	XactStart *string `json:"xact_start,omitempty"`

	// **参数解释**: 事务持续时间，单位秒。 **取值范围**: 不涉及。
	TransactionTime *string `json:"transaction_time,omitempty"`

	// **参数解释**: 查询开始时间。 **取值范围**: 不涉及。
	QueryStart *string `json:"query_start,omitempty"`

	// **参数解释**: 上次状态改变的时间。 **取值范围**: 不涉及。
	StateChange *string `json:"state_change,omitempty"`

	// **参数解释**: 查询持续时间，单位秒。 **取值范围**: 不涉及。
	QueryTime *string `json:"query_time,omitempty"`

	// **参数解释**: 会话开始时间。 **取值范围**: 不涉及。
	BackendStart *string `json:"backend_start,omitempty"`

	// **参数解释**: 是否在等待锁。 **取值范围**: 不涉及。
	Waiting *string `json:"waiting,omitempty"`

	// **参数解释**: （等待获取的）锁模式。 **取值范围**: 不涉及。
	Lockmode *string `json:"lockmode,omitempty"`

	// **参数解释**: 阻塞会话ID。 **取值范围**: 不涉及。
	BlockSessionid *string `json:"block_sessionid,omitempty"`

	// **参数解释**: 阻塞会话数。 **取值范围**: 不涉及。
	BlockCount *string `json:"block_count,omitempty"`

	// **参数解释**: 归一化SQL ID。 **取值范围**: 不涉及。
	UniqueSqlId *string `json:"unique_sql_id,omitempty"`

	// **参数解释**: 查询 ID。 **取值范围**: 不涉及。
	QueryId *string `json:"query_id,omitempty"`

	// **参数解释**: SQL文本。 **取值范围**: 不涉及。
	Query *string `json:"query,omitempty"`

	// **参数解释**: 当前事务ID。 **取值范围**: 不涉及。
	CurrentXid *string `json:"current_xid,omitempty"`

	// **参数解释**: 顶层事务ID。 **取值范围**: 不涉及。
	TopXid *string `json:"top_xid,omitempty"`

	// **参数解释**: 事务当前使用的XLOG量，单位为字节。 **取值范围**: 不涉及。
	XlogQuantity *string `json:"xlog_quantity,omitempty"`

	// **参数解释**: 该会话当前总体状态。 **取值范围**: -active：后台正在执行一个查询。 -idle：后台正在等待一个新的客户端命令。 -idle in transaction：后台在事务中，但事务中没有语句在执行。 -fastpath function call：后台正在执行一个fast-path函数。 -disabled：如果后台禁用track_activities，则事务显示此状态。
	State *string `json:"state,omitempty"`

	// **参数解释**: 应用名称。 **取值范围**: 不涉及。
	ApplicationName *string `json:"application_name,omitempty"`

	// **参数解释**: 连接信息。 **取值范围**: 不涉及。
	ConnectionInfo *string `json:"connection_info,omitempty"`
}

func (o WaitEventResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WaitEventResult struct{}"
	}

	return strings.Join([]string{"WaitEventResult", string(data)}, " ")
}
