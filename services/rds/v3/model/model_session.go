package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Session 会话信息
type Session struct {

	// 采样时间
	SampleTime *string `json:"sample_time,omitempty"`

	// 阻塞进程ID
	BlockedProcessId *string `json:"blocked_process_id,omitempty"`

	// 数据库OID
	DatabaseOid *int32 `json:"database_oid,omitempty"`

	// 数据库名
	DatabaseName *string `json:"database_name,omitempty"`

	// 会话ID
	SessionId *int32 `json:"session_id,omitempty"`

	// 并行会话ID
	ParallelLeaderId *int32 `json:"parallel_leader_id,omitempty"`

	// 后端用户OID
	BackendUserOid *int32 `json:"backend_user_oid,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 应用名
	AppName *string `json:"app_name,omitempty"`

	// 客户端地址
	ClientIpAddress *string `json:"client_ip_address,omitempty"`

	// 客户端名称
	ClientHostName *string `json:"client_host_name,omitempty"`

	// 客户端端口
	ClientPort *int32 `json:"client_port,omitempty"`

	// 会话建立时间
	SessionStartTime *string `json:"session_start_time,omitempty"`

	// 事务启动时间
	TransactionStartTime *string `json:"transaction_start_time,omitempty"`

	// 事务执行时间(s)
	TransactionExecutionTime *int32 `json:"transaction_execution_time,omitempty"`

	// 查询开始时间
	QueryStartTime *string `json:"query_start_time,omitempty"`

	// state改变时间
	StateChangeTime *string `json:"state_change_time,omitempty"`

	// 等待事件类型
	WaitEventType *string `json:"wait_event_type,omitempty"`

	// 等待事件名称
	WaitEventName *string `json:"wait_event_name,omitempty"`

	// 会话状态
	SessionStatus *string `json:"session_status,omitempty"`

	// Backend XID
	BackendXid *int32 `json:"backend_xid,omitempty"`

	// Backend Xmin
	BackendXmin *int32 `json:"backend_xmin,omitempty"`

	// Query ID
	QueryId *string `json:"query_id,omitempty"`

	// SQL语句
	SqlStatement *string `json:"sql_statement,omitempty"`

	// 进程类型
	ProcessType *string `json:"process_type,omitempty"`

	// 内存占比(%)
	MemoryUsage *float32 `json:"memory_usage,omitempty"`

	// 进程状态
	ProcessStatus *string `json:"process_status,omitempty"`

	// 3秒内平均CPU占用率(%)
	CpuUsage *float32 `json:"cpu_usage,omitempty"`

	// I/O等待时间(s)
	IoWaitStatus *float32 `json:"io_wait_status,omitempty"`

	// 磁盘读速率(MB/s)
	DiskReadRate *float32 `json:"disk_read_rate,omitempty"`

	// 磁盘写速率(MB/s)
	DiskWriteRate *float32 `json:"disk_write_rate,omitempty"`
}

func (o Session) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Session struct{}"
	}

	return strings.Join([]string{"Session", string(data)}, " ")
}
