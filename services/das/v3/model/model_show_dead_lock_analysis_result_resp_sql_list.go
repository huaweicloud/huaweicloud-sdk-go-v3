package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowDeadLockAnalysisResultRespSqlList struct {

	// SQL语句
	Sql string `json:"sql"`

	// 发生时间
	OccurrenceTime int64 `json:"occurrence_time"`

	// 执行耗时毫秒
	QueryTime int64 `json:"query_time"`

	// 事务ID
	TransactionId string `json:"transaction_id"`

	// 模板ID
	SqlTemplateId string `json:"sql_template_id"`

	// 节点ID
	NodeId string `json:"node_id"`

	// 用户名
	DbUser string `json:"db_user"`

	// 数据库
	Database string `json:"database"`

	// 客户端IP
	Client string `json:"client"`

	// SQL类型
	SqlType string `json:"sql_type"`

	// 执行状态
	Status int64 `json:"status"`

	// 错误码
	ErrorNo int64 `json:"error_no"`

	// 更新行数
	RowsAffected int64 `json:"rows_affected"`

	// 返回行数
	RowsSent int64 `json:"rows_sent"`

	// 锁等待时间毫秒
	LockTime int64 `json:"lock_time"`

	// 扫描行数
	RowsExamined int64 `json:"rows_examined"`

	// 线程ID
	SessionId string `json:"session_id"`

	// CPU耗时(us)
	CpuTime int64 `json:"cpu_time"`

	// 网络发送字节数
	SendBytes int64 `json:"send_bytes"`

	// 查询中所有的表名
	QueryTables string `json:"query_tables"`

	// 物理I/O读字节数
	InnodbIoReadBytes int64 `json:"innodb_io_read_bytes"`

	// 物理I/O读次数
	InnodbIoRead int64 `json:"innodb_io_read"`

	// 物理I/O读取等待耗时(ms)
	InnodbIoReadWait int64 `json:"innodb_io_read_wait"`

	// 物理I/O读取等待耗时(ms)
	InnodbLockWait int64 `json:"innodb_lock_wait"`

	// 行锁等待耗时(ms)
	InnodbQueueWait int64 `json:"innodb_queue_wait"`
}

func (o ShowDeadLockAnalysisResultRespSqlList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeadLockAnalysisResultRespSqlList struct{}"
	}

	return strings.Join([]string{"ShowDeadLockAnalysisResultRespSqlList", string(data)}, " ")
}
