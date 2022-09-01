package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InnodbTrx struct {

	// 事务ID
	TrxId string `json:"trx_id" xml:"trx_id"`

	// 事务状态
	TrxState string `json:"trx_state" xml:"trx_state"`

	// 事务开始时间
	TrxStarted string `json:"trx_started" xml:"trx_started"`

	// 事务等待开始时间
	TrxWaitStarted string `json:"trx_wait_started" xml:"trx_wait_started"`

	// 会话ID，同ListProcesses接口返回的id。
	TrxMysqlThreadId string `json:"trx_mysql_thread_id" xml:"trx_mysql_thread_id"`

	// 事务运行的SQL语句
	TrxQuery string `json:"trx_query" xml:"trx_query"`

	// 加行锁的表数量
	TrxTablesLocked string `json:"trx_tables_locked" xml:"trx_tables_locked"`

	// 锁定的行数量（近似值）
	TrxRowsLocked string `json:"trx_rows_locked" xml:"trx_rows_locked"`

	// 事务插入或者修改的行数
	TrxRowsModified string `json:"trx_rows_modified" xml:"trx_rows_modified"`

	// 隔离级别
	TrxIsolationLevel string `json:"trx_isolation_level" xml:"trx_isolation_level"`

	// 等待锁信息
	InnodbWaitLocks []InnodbLock `json:"innodb_wait_locks" xml:"innodb_wait_locks"`

	// 持有锁信息
	InnodbHoldLocks []InnodbLock `json:"innodb_hold_locks" xml:"innodb_hold_locks"`
}

func (o InnodbTrx) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InnodbTrx struct{}"
	}

	return strings.Join([]string{"InnodbTrx", string(data)}, " ")
}
