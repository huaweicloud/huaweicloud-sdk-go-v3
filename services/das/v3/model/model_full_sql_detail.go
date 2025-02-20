package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FullSqlDetail struct {

	// SQL语句。
	Sql string `json:"sql"`

	// SQL模板ID。
	SqlTemplateId string `json:"sql_template_id"`

	// 操作类型。
	OperateType string `json:"operate_type"`

	// 状态。
	Status string `json:"status"`

	// 错误码。
	ErrorNo string `json:"error_no"`

	// 数据库。
	Database string `json:"database"`

	// 线程ID。
	ThreadId string `json:"thread_id"`

	// 客户端。
	Client string `json:"client"`

	// 用户。
	User string `json:"user"`

	// 执行开始时间（Unix timestamp），单位：毫秒。
	ExecuteAt int64 `json:"execute_at"`

	// 执行耗时（毫秒）。
	QueryTime float64 `json:"query_time"`

	// 锁等待耗时（毫秒）。
	LockTime float64 `json:"lock_time"`

	// 扫描行数。
	RowsExamined int64 `json:"rows_examined"`

	// 返回行数。
	RowsSent int64 `json:"rows_sent"`

	// 更新行数。
	RowsAffected int64 `json:"rows_affected"`

	// 事务ID。
	TrxId int64 `json:"trx_id"`

	// CPU耗时（微秒）。
	CpuTime float64 `json:"cpu_time"`

	// 网络发送字节数。
	SendBytes int64 `json:"send_bytes"`

	// 查询中所有的表名（格式：库名.表名|库名.表名）。
	QueryTables string `json:"query_tables"`

	// 物理I/O读字节数。
	InnodbIoReadBytes int64 `json:"innodb_io_read_bytes"`

	// 物理I/O读次数。
	InnodbIoRead int32 `json:"innodb_io_read"`

	// 物理I/O读取等待耗时（毫秒）。
	InnodbIoReadWait float64 `json:"innodb_io_read_wait"`

	// 行锁等待耗时（毫秒）。
	InnodbLockWait float64 `json:"innodb_lock_wait"`

	// 进入innodb的等待耗时（毫秒）。
	InnodbQueueWait float64 `json:"innodb_queue_wait"`

	// 内核版本号，DDM实例使用。
	KernelVersion string `json:"kernel_version"`

	// SQL执行各阶段细分耗时，DDM实例使用。
	QueryTimeDetail string `json:"query_time_detail"`

	// 会话ID。
	SessionId string `json:"session_id"`

	// 节点ID。
	NodeId string `json:"node_id"`
}

func (o FullSqlDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullSqlDetail struct{}"
	}

	return strings.Join([]string{"FullSqlDetail", string(data)}, " ")
}
