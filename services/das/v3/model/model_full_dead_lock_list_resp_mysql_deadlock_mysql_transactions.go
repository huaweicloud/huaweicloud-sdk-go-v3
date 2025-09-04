package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FullDeadLockListRespMysqlDeadlockMysqlTransactions struct {

	// 表
	Table *string `json:"table,omitempty"`

	// 具体sql
	Sql *string `json:"sql,omitempty"`

	// 会话id
	SessionId *string `json:"session_id,omitempty"`

	// 线程id
	ThreadId *string `json:"thread_id,omitempty"`

	// 请求类型
	RequestType *string `json:"request_type,omitempty"`

	// 内部事务id
	TransactionId *string `json:"transaction_id,omitempty"`

	// 事务正在等待的锁的描述
	WaitingLock *string `json:"waiting_lock,omitempty"`

	// 锁针对的索引
	WaitingLockIndex *string `json:"waiting_lock_index,omitempty"`

	// 锁类型
	WaitingLockType *string `json:"waiting_lock_type,omitempty"`
}

func (o FullDeadLockListRespMysqlDeadlockMysqlTransactions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullDeadLockListRespMysqlDeadlockMysqlTransactions struct{}"
	}

	return strings.Join([]string{"FullDeadLockListRespMysqlDeadlockMysqlTransactions", string(data)}, " ")
}
