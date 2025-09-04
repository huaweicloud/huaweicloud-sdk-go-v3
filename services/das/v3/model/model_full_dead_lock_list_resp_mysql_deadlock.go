package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FullDeadLockListRespMysqlDeadlock 死锁事件详细信息
type FullDeadLockListRespMysqlDeadlock struct {

	// 死锁事件详细信息
	Raw *string `json:"raw,omitempty"`

	// 死锁事件发生的 Unix 时间戳（毫秒级）
	HappenTime *int64 `json:"happen_time,omitempty"`

	// MySQL InnoDB 引擎在检测到死锁后，会自动选择一个事务进行回滚以解除死锁。
	RollbackTrxId *string `json:"rollback_trx_id,omitempty"`

	// 参与这次死锁的每个事务的详细结构化信息
	MysqlTransactions *[]FullDeadLockListRespMysqlDeadlockMysqlTransactions `json:"mysql_transactions,omitempty"`
}

func (o FullDeadLockListRespMysqlDeadlock) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullDeadLockListRespMysqlDeadlock struct{}"
	}

	return strings.Join([]string{"FullDeadLockListRespMysqlDeadlock", string(data)}, " ")
}
