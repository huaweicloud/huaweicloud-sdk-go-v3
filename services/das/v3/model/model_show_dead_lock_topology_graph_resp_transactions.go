package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowDeadLockTopologyGraphRespTransactions struct {

	// 事务ID
	TransactionId string `json:"transaction_id"`

	// 线程ID
	ThreadId int64 `json:"thread_id"`

	// 是否被回滚
	RollbackTarget bool `json:"rollback_target"`

	// SQL语句
	Sql string `json:"sql"`

	// 操作的表名
	Table string `json:"table"`

	// 操作类型
	Operator string `json:"operator"`

	// 行锁数量
	RowLockCount int64 `json:"row_lock_count"`

	// Undo日志条数
	UndoLogEntries int64 `json:"undo_log_entries"`
}

func (o ShowDeadLockTopologyGraphRespTransactions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeadLockTopologyGraphRespTransactions struct{}"
	}

	return strings.Join([]string{"ShowDeadLockTopologyGraphRespTransactions", string(data)}, " ")
}
