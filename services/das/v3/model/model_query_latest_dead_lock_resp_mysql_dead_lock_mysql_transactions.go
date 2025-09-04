package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryLatestDeadLockRespMysqlDeadLockMysqlTransactions struct {

	// 事务操作的具体表
	Table *string `json:"table,omitempty"`

	// 前正在执行或最近执行的 SQL 语句
	Sql *string `json:"sql,omitempty"`

	// 会话的 ID
	SessionId *string `json:"session_id,omitempty"`

	// 线程id
	ThreadId *string `json:"thread_id,omitempty"`

	// 请求类型
	RequestType *string `json:"request_type,omitempty"`

	// 内部的事务 ID
	TransactionId *string `json:"transaction_id,omitempty"`

	// 事务正在等待的锁的详细描述
	WaitingLock *string `json:"waiting_lock,omitempty"`

	// 正在等待的锁是针对表的索引
	WaitingLockIndex *string `json:"waiting_lock_index,omitempty"`

	// 正在等待的锁的具体类型
	WaitingLockType *string `json:"waiting_lock_type,omitempty"`

	// 该事务当前持有的锁的详细描述
	HoldingLock *string `json:"holding_lock,omitempty"`

	// 持有的锁的针对表的索引
	HoldingLockIndex *string `json:"holding_lock_index,omitempty"`

	// 持有的锁的具体类型。X 是排他锁，rec 是记录锁
	HoldingLockType *string `json:"holding_lock_type,omitempty"`
}

func (o QueryLatestDeadLockRespMysqlDeadLockMysqlTransactions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryLatestDeadLockRespMysqlDeadLockMysqlTransactions struct{}"
	}

	return strings.Join([]string{"QueryLatestDeadLockRespMysqlDeadLockMysqlTransactions", string(data)}, " ")
}
