package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SqlItemDto SQL项详情
type SqlItemDto struct {

	// SQL的ID值
	Id *string `json:"id,omitempty"`

	// 操作类型
	OperateType *string `json:"operate_type,omitempty"`

	// 模板ID
	SqlTemplateId *string `json:"sql_template_id,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// SQL文本
	Sql *string `json:"sql,omitempty"`

	// 数据库名称
	Database *string `json:"database,omitempty"`

	// 线程ID
	ThreadId *int64 `json:"thread_id,omitempty"`

	// 用户名称
	Username *string `json:"username,omitempty"`

	// 客户端IP
	ClientIp *string `json:"client_ip,omitempty"`

	// 执行状态
	Status *int32 `json:"status,omitempty"`

	// 执行耗时(ms)
	ExecuteCost *float64 `json:"execute_cost,omitempty"`

	// 执行时间点(ms)
	ExecuteAt *float64 `json:"execute_at,omitempty"`

	// 更新行数
	RowsAffected *int32 `json:"rows_affected,omitempty"`

	// 扫描行数
	RowsExamined *int32 `json:"rows_examined,omitempty"`

	// 锁等待时间
	LockWaitTime *float64 `json:"lock_wait_time,omitempty"`

	// 返回行数
	RowsReturned *int32 `json:"rows_returned,omitempty"`

	// 事务ID
	TrxId *int64 `json:"trx_id,omitempty"`

	// CPU耗时
	CpuTime *int32 `json:"cpu_time,omitempty"`

	// 网络发送字节数
	SendBytes *int64 `json:"send_bytes,omitempty"`

	// 查询中所有的表名（格式：库名.表名|库名.表名）
	QueryTables *string `json:"query_tables,omitempty"`

	// 物理IO读字节数
	InnodbIoReadBytes *int64 `json:"innodb_io_read_bytes,omitempty"`

	// 物理IO读次数
	InnodbIoRead *int32 `json:"innodb_io_read,omitempty"`

	// 物理IO读取等待耗时（ms）
	InnodbIoReadWait *float64 `json:"innodb_io_read_wait,omitempty"`

	// 行锁等待耗时（ms）
	InnodbLockWait *float64 `json:"innodb_lock_wait,omitempty"`

	// 进入innodb的等待耗时（ms）
	InnodbQueueWait *float64 `json:"innodb_queue_wait,omitempty"`

	// 内核版本号
	KernelVersion *string `json:"kernel_version,omitempty"`

	// SQL执行各阶段细分耗时
	QueryTimeDetail *string `json:"query_time_detail,omitempty"`

	// 会话ID
	SessionId *string `json:"session_id,omitempty"`

	// 错误码
	ErrorNo *int32 `json:"error_no,omitempty"`

	// 节点ID，实例节点的唯一标识
	NodeId *string `json:"node_id,omitempty"`

	// sqlserver IO逻辑读
	LogicalReads *int64 `json:"logical_reads,omitempty"`

	// sqlserver IO物理读
	PhysicalReads *int64 `json:"physical_reads,omitempty"`

	// sqlserver IO写
	Writes *int64 `json:"writes,omitempty"`

	// sqlserver 应用名
	AppName *string `json:"app_name,omitempty"`
}

func (o SqlItemDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlItemDto struct{}"
	}

	return strings.Join([]string{"SqlItemDto", string(data)}, " ")
}
