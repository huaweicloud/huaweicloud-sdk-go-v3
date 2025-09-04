package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryLatestDeadLockRespMysqlDeadLock MySQL 死锁事件的所有详细信息
type QueryLatestDeadLockRespMysqlDeadLock struct {

	// 原始的、未经解析的 MySQL 死锁日志文本
	Raw *string `json:"raw,omitempty"`

	// 死锁事件发生的日期和时间
	Time *string `json:"time,omitempty"`

	// 死锁事件发生的 Unix 时间戳（毫秒级）
	HappenTime *int64 `json:"happen_time,omitempty"`

	// MySQL 在检测到死锁后，会自动选择一个事务进行回滚（通常是代价较小的那个）
	RollbackTrxId *string `json:"rollback_trx_id,omitempty"`

	// 参与死锁的每个事务的详细信息
	MysqlTransactions *[]QueryLatestDeadLockRespMysqlDeadLockMysqlTransactions `json:"mysql_transactions,omitempty"`
}

func (o QueryLatestDeadLockRespMysqlDeadLock) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryLatestDeadLockRespMysqlDeadLock struct{}"
	}

	return strings.Join([]string{"QueryLatestDeadLockRespMysqlDeadLock", string(data)}, " ")
}
