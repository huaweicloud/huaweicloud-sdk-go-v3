package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MySqlTransaction MySQL事务信息
type MySqlTransaction struct {

	// 会话ID
	SessionId *string `json:"session_id,omitempty"`

	// 线程ID
	ThreadId *string `json:"thread_id,omitempty"`

	// 请求类型
	RequestType *string `json:"request_type,omitempty"`

	// 事务ID
	TransactionId *string `json:"transaction_id,omitempty"`

	// 涉及表
	Table *string `json:"table,omitempty"`

	// 等待锁
	WaitingLock *string `json:"waiting_lock,omitempty"`

	// 等待锁索引名
	WaitingLockIndex *string `json:"waiting_lock_index,omitempty"`

	// 等待锁索引类型
	WaitingLockType *string `json:"waiting_lock_type,omitempty"`

	// 持有锁
	HoldingLock *string `json:"holding_lock,omitempty"`

	// 持有锁索引
	HoldingLockIndex *string `json:"holding_lock_index,omitempty"`

	// 持有锁索引类型
	HoldingLockType *string `json:"holding_lock_type,omitempty"`

	// SQL语句
	Sql *string `json:"sql,omitempty"`
}

func (o MySqlTransaction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MySqlTransaction struct{}"
	}

	return strings.Join([]string{"MySqlTransaction", string(data)}, " ")
}
