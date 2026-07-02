package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LockRootCause struct {

	// **参数解释**： 被锁会话ID。 **取值范围**： 不涉及。
	LockedPid *int64 `json:"locked_pid,omitempty"`

	// **参数解释**： 被锁会话当前执行的SQL。 **取值范围**： 不涉及。
	LockedSql *string `json:"locked_sql,omitempty"`

	// **参数解释**： 等待持续时间（秒）。 **取值范围**： 不涉及。
	WaitSeconds *int64 `json:"wait_seconds,omitempty"`

	// **参数解释**： 锁所在数据库。 **取值范围**： 不涉及。
	LockDb *string `json:"lock_db,omitempty"`

	// **参数解释**： 锁所在表。 **取值范围**： 不涉及。
	LockTable *string `json:"lock_table,omitempty"`

	// **参数解释**： 锁所在索引。 **取值范围**： 不涉及。
	LockIndex *string `json:"lock_index,omitempty"`

	// **参数解释**： 锁模式。 **取值范围**： - IX：表级排他锁。 - X,REC_NOT_GAP：排他记录锁。 - X,GAP：间隙锁。 - X：行级排他锁。
	LockMode *string `json:"lock_mode,omitempty"`

	// **参数解释**： 锁住的具体数据行的标识。 **取值范围**： 不涉及。
	LockData *string `json:"lock_data,omitempty"`

	// **参数解释**： 阻塞源会话ID。 **取值范围**： 不涉及。
	BlockerPid *int64 `json:"blocker_pid,omitempty"`

	// **参数解释**： 阻塞源事务状态。 **取值范围**： - RUNNING：运行中。 - LOCK WAIT：锁等待。 - ROLLING BACK：回滚中。 - COMMITTING：提交中。
	BlockerState *string `json:"blocker_state,omitempty"`

	// **参数解释**： 阻塞源事务持续时间（秒）。 **取值范围**： 不涉及。
	BlockerAge *int32 `json:"blocker_age,omitempty"`

	// **参数解释**： 阻塞源锁定的行数。 **取值范围**： 不涉及。
	BlockerRowsLocked *int64 `json:"blocker_rows_locked,omitempty"`

	// **参数解释**： 阻塞源修改的行数。 **取值范围**： 不涉及。
	BlockerRowsModified *int64 `json:"blocker_rows_modified,omitempty"`

	// **参数解释**： 阻塞源当前执行的SQL列表。
	BlockerCurrentSql *[]string `json:"blocker_current_sql,omitempty"`

	// **参数解释**： 阻塞源主机。 **取值范围**： 不涉及。
	BlockerHost *string `json:"blocker_host,omitempty"`

	// **参数解释**： 阻塞源命令。 **取值范围**： 不涉及。
	BlockerCommand *string `json:"blocker_command,omitempty"`

	// **参数解释**： 阻塞源线程ID。 **取值范围**： 不涉及。
	BlockerThreadId *int64 `json:"blocker_thread_id,omitempty"`
}

func (o LockRootCause) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LockRootCause struct{}"
	}

	return strings.Join([]string{"LockRootCause", string(data)}, " ")
}
