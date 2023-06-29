package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetadataLock 元数据锁
type MetadataLock struct {

	// 会话ID
	ThreadId string `json:"thread_id"`

	// 锁状态，取值为PENDING和GRANTED，分别表示等待锁和持有锁。
	LockStatus string `json:"lock_status"`

	// 加锁模式，取值为MDL_SHARED 、MDL_EXCLUSIVE 、MDL_SHARED_READ、MDL_SHARED_WRITE等。
	LockMode string `json:"lock_mode"`

	// 锁类型，取值为Table metadata lock、Schema metadata lock、Tablespace lock、Global read lock，分别表示表元数据锁、库元数据锁、表空间锁、全局读锁。
	LockType string `json:"lock_type"`

	// 锁范围，取值为MDL_STATEMENT、MDL_TRANSACTION、MDL_EXPLICIT，分别表示语句级别、事务级别、global级别
	LockDuration string `json:"lock_duration"`

	// 锁所在的数据库，对于部分Global read lock级别的元数据锁，该值为空。
	TableSchema string `json:"table_schema"`

	// 表名
	TableName string `json:"table_name"`

	// 用户
	User string `json:"user"`

	// 时间
	Time string `json:"time"`

	// 主机
	Host string `json:"host"`

	// 会话所在的数据库
	Database string `json:"database"`

	// 命令
	Command string `json:"command"`

	// 状态
	State string `json:"state"`

	// SQL语句
	Sql string `json:"sql"`

	// 事务执行时间
	TrxExecTime string `json:"trx_exec_time"`

	// 阻塞会话列表
	BlockProcess []Process `json:"block_process"`

	// 等待会话列表
	WaitProcess []Process `json:"wait_process"`
}

func (o MetadataLock) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetadataLock struct{}"
	}

	return strings.Join([]string{"MetadataLock", string(data)}, " ")
}
