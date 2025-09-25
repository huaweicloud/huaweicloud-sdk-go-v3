package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListTransactionResponseBodyRowsInfo struct {

	// **参数解释**: 事务ID。 **取值范围**: 不涉及。
	Sessionid *int32 `json:"sessionid,omitempty"`

	// **参数解释**: 线程ID。 **取值范围**: 不涉及。
	Pid *int32 `json:"pid,omitempty"`

	// **参数解释**: SQL查询ID。 **取值范围**: 不涉及。
	QueryId *int32 `json:"query_id,omitempty"`

	// **参数解释**: 数据库。 **取值范围**: 不涉及。
	Datname *string `json:"datname,omitempty"`

	// **参数解释**: 用户发起事务请求的客户端地址。 **取值范围**: 不涉及。
	ClientAddr *string `json:"client_addr,omitempty"`

	// **参数解释**: 用户发起事务请求的客户端端口。 **取值范围**: 不涉及。
	ClientPort *int32 `json:"client_port,omitempty"`

	// **参数解释**: 用户名。 **取值范围**: 不涉及。 **取值范围**: 不涉及。
	Usename *string `json:"usename,omitempty"`

	// **参数解释**: 查询的SQL语句。 **取值范围**: 不涉及。
	Query *string `json:"query,omitempty"`

	// **参数解释**: 会话开始时间。 **取值范围**: 不涉及。
	BackendStart *string `json:"backend_start,omitempty"`

	// **参数解释**: 事务开始时间。 **取值范围**: 不涉及。
	XactStart *string `json:"xact_start,omitempty"`

	// **参数解释**: 应用名称。 **取值范围**: 不涉及。
	ApplicationName *string `json:"application_name,omitempty"`

	// **参数解释**: 事务状态。 **取值范围**: 不涉及。
	State *string `json:"state,omitempty"`

	// **参数解释**: 事务变更时间。 **取值范围**: 不涉及。
	StateChange *string `json:"state_change,omitempty"`

	// **参数解释**: 查询开始时间。 **取值范围**: 不涉及。
	QueryStart *string `json:"query_start,omitempty"`

	// **参数解释**: 是否等待。 **取值范围**: 不涉及。
	Waiting *string `json:"waiting,omitempty"`

	// **参数解释**: 归一化ID。 **取值范围**: 不涉及。
	UniqueSqlId *int32 `json:"unique_sql_id,omitempty"`

	// **参数解释**: 顶层事务ID。
	TopXid *string `json:"top_xid,omitempty"`

	// **参数解释**: 当前事务ID。 **取值范围**: 不涉及。
	CurrentXid *string `json:"current_xid,omitempty"`

	// **参数解释**: 事务执行时长。 **取值范围**: 不涉及。
	ExecTime *string `json:"exec_time,omitempty"`

	// **参数解释**: 事务xlog量。 **取值范围**: 不涉及。
	XlogQuantity *int32 `json:"xlog_quantity,omitempty"`
}

func (o ListTransactionResponseBodyRowsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransactionResponseBodyRowsInfo struct{}"
	}

	return strings.Join([]string{"ListTransactionResponseBodyRowsInfo", string(data)}, " ")
}
