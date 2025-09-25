package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WaitEventQueryInfo struct {

	// **参数解释**: 数据库名称。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	DatabaseName *string `json:"database_name,omitempty"`

	// **参数解释**: 用户名称。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: 是否在等待状态。 **约束限制**: 不涉及。 **取值范围**: -t：是。 -f：否。 **默认取值**: 不涉及。
	Waiting *string `json:"waiting,omitempty"`

	// **参数解释**: 会话ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SessionId *string `json:"session_id,omitempty"`

	// **参数解释**: 阻塞当前会话的会话ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	BlockSessionId *string `json:"block_session_id,omitempty"`

	// **参数解释**: 阻塞当前会话的会话数。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	BlockCount *string `json:"block_count,omitempty"`

	// **参数解释**: 唯一的SQL ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	UniqueSqlId *string `json:"unique_sql_id,omitempty"`

	// **参数解释**: SQL查询ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	QueryId *string `json:"query_id,omitempty"`

	// **参数解释**: 该会话当前整体状态。 **约束限制**: 不涉及。 **取值范围**: -active：后台正在执行一个查询。 -idle：后台正在等待一个新的客户端命令。 -idle in transaction：后台在事务中，但事务中没有语句在执行。 -fastpath function call：后台正在执行一个fast-path函数。 -disabled：如果后台禁用track_activities，则事务显示此状态。 **默认取值**: 不涉及。
	State *string `json:"state,omitempty"`

	// **参数解释**: 等待事件。 参见“开发指南 > 系统表和系统视图 > 系统视图 > 其他系统视图 > PG_THREAD_WAIT_STATUS”中的wait_event字段。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	WaitEvent *string `json:"wait_event,omitempty"`

	// **参数解释**: 等待状态。 参见“开发指南 > 系统表和系统视图 > 系统视图 > 其他系统视图 > PG_THREAD_WAIT_STATUS”中的wait_status列表。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	WaitStatus *string `json:"wait_status,omitempty"`
}

func (o WaitEventQueryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WaitEventQueryInfo struct{}"
	}

	return strings.Join([]string{"WaitEventQueryInfo", string(data)}, " ")
}
