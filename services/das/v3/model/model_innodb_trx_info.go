package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InnodbTrxInfo InnoDB事务信息
type InnodbTrxInfo struct {

	// 事务ID
	TrxId *string `json:"trx_id,omitempty"`

	// 事务状态
	TrxState *string `json:"trx_state,omitempty"`

	// 事务开始时间
	TrxStarted *string `json:"trx_started,omitempty"`

	// 事务开始时间戳
	TrxStartedTimestamp *int64 `json:"trx_started_timestamp,omitempty"`

	// 事务当前正在等待锁的Id
	TrxRequestedLockId *string `json:"trx_requested_lock_id,omitempty"`

	// 事务开始等待时间
	TrxWaitStarted *string `json:"trx_wait_started,omitempty"`

	// 事务开始等待时间戳
	TrxWaitStartedTimestamp *int64 `json:"trx_wait_started_timestamp,omitempty"`

	// 事务权重
	TrxWeight *string `json:"trx_weight,omitempty"`

	// 会话ID
	TrxMysqlThreadId *string `json:"trx_mysql_thread_id,omitempty"`

	// 事务正在执行的SQL语句
	TrxQuery *string `json:"trx_query,omitempty"`

	// 事务当前操作状态
	TrxOperationState *string `json:"trx_operation_state,omitempty"`

	// 当前事务执行的SQL中使用的表个数
	TrxTablesInUse *string `json:"trx_tables_in_use,omitempty"`

	// 当前执行SQL的行锁数量
	TrxTablesLocked *string `json:"trx_tables_locked,omitempty"`

	// 事务保留的锁数量
	TrxLockStructs *string `json:"trx_lock_structs,omitempty"`

	// 事务锁住的内存大小
	TrxLockMemoryBytes *string `json:"trx_lock_memory_bytes,omitempty"`

	// 事务锁住的行记录数
	TrxRowsLocked *string `json:"trx_rows_locked,omitempty"`

	// 事务更改的行数
	TrxRowsModified *string `json:"trx_rows_modified,omitempty"`

	// 事务并发票数
	TrxConcurrencyTickets *string `json:"trx_concurrency_tickets,omitempty"`

	// 事务隔离级别
	TrxIsolationLevel *string `json:"trx_isolation_level,omitempty"`
}

func (o InnodbTrxInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InnodbTrxInfo struct{}"
	}

	return strings.Join([]string{"InnodbTrxInfo", string(data)}, " ")
}
