package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MySqlLatestDeadLock MySQL死锁内容
type MySqlLatestDeadLock struct {

	// 原始死锁内容
	Raw *string `json:"raw,omitempty"`

	// 发生时间（ms）
	HappenTime *int64 `json:"happen_time,omitempty"`

	// 死锁时间
	Time *string `json:"time,omitempty"`

	// 回滚事务ID
	RollbackTrxId *string `json:"rollback_trx_id,omitempty"`

	// 事务列表
	MysqlTransactions *[]MySqlTransaction `json:"mysql_transactions,omitempty"`
}

func (o MySqlLatestDeadLock) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MySqlLatestDeadLock struct{}"
	}

	return strings.Join([]string{"MySqlLatestDeadLock", string(data)}, " ")
}
