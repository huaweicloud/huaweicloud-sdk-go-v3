package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MySqlDeadLock MySQL死锁内容
type MySqlDeadLock struct {

	// 原始死锁内容
	Raw *string `json:"raw,omitempty"`

	// 发生时间（ms）
	HappenTime *int64 `json:"happen_time,omitempty"`

	// 回滚事务ID
	RollbackTrxId *string `json:"rollback_trx_id,omitempty"`

	// 事务列表
	MysqlTransactions *[]MySqlTransaction `json:"mysql_transactions,omitempty"`
}

func (o MySqlDeadLock) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MySqlDeadLock struct{}"
	}

	return strings.Join([]string{"MySqlDeadLock", string(data)}, " ")
}
