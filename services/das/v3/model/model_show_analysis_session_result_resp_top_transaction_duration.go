package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowAnalysisSessionResultRespTopTransactionDuration struct {

	// 会话ID
	SessionId string `json:"session_id"`

	// 用户名
	User string `json:"user"`

	// 主机IP
	Host string `json:"host"`

	// 数据库名
	DatabaseName string `json:"database_name"`

	// 执行状态
	ExecutionStatus string `json:"execution_status"`

	// 命令类型
	Command string `json:"command"`

	// SQL语句
	SqlStatement string `json:"sql_statement"`

	// 状态持续时间（秒）
	StateDuration string `json:"state_duration"`

	// 事务持续时间（秒）
	TransactionDuration string `json:"transaction_duration"`

	// 事务ID
	TransactionId string `json:"transaction_id"`

	// 事务锁等待时长（秒）
	TransactionLockWaitTime string `json:"transaction_lock_wait_time"`

	// 事务状态
	TransactionStatus string `json:"transaction_status"`

	// 事务锁定行数
	RowsLockedByTransactions string `json:"rows_locked_by_transactions"`

	// 事务锁定表数量
	TablesLockedByTransactions string `json:"tables_locked_by_transactions"`

	// 事务更新行数
	RowsUpdatedByTransactions string `json:"rows_updated_by_transactions"`
}

func (o ShowAnalysisSessionResultRespTopTransactionDuration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAnalysisSessionResultRespTopTransactionDuration struct{}"
	}

	return strings.Join([]string{"ShowAnalysisSessionResultRespTopTransactionDuration", string(data)}, " ")
}
