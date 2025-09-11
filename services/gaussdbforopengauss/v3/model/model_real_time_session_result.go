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
}

func (o RealTimeSessionResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RealTimeSessionResult struct{}"
	}

	return strings.Join([]string{"RealTimeSessionResult", string(data)}, " ")
}
