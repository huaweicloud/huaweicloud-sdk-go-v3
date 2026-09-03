package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetaLockInfo MDL锁等待信息
type MetaLockInfo struct {

	// MDL锁ID
	LockId *string `json:"lock_id,omitempty"`

	// 线程ID
	ThreadId *string `json:"thread_id,omitempty"`

	// MDL锁状态
	LockStatus *string `json:"lock_status,omitempty"`

	// MDL锁等待模式
	LockMode *string `json:"lock_mode,omitempty"`

	// MDL锁等待信息
	LockType *string `json:"lock_type,omitempty"`

	// MDL锁等待持续时间
	LockDuration *string `json:"lock_duration,omitempty"`

	// 库表schema信息
	TableSchema *string `json:"table_schema,omitempty"`

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// 用户名称
	User *string `json:"user,omitempty"`

	// MDL锁等待时间
	Time *string `json:"time,omitempty"`

	// MDL锁等待阻塞数量
	BlockNumber *int32 `json:"block_number,omitempty"`

	// MDL锁等待数量
	WaitNumber *int32 `json:"wait_number,omitempty"`

	// 主机
	Host *string `json:"host,omitempty"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`

	// MDL锁等待SQL语句
	Command *string `json:"command,omitempty"`

	// MDL锁等待状态
	State *string `json:"state,omitempty"`

	// MDL锁等待额外信息
	Info *string `json:"info,omitempty"`

	// 关联的SQL限流规则
	SqlLimitRule *string `json:"sql_limit_rule,omitempty"`

	// 事务执行时间
	TrxExecTime *string `json:"trx_exec_time,omitempty"`

	// 阻塞的事务信息列表
	BlockProcessInfo *[]ProcessInfo `json:"block_process_info,omitempty"`

	// 等待的事务信息列表
	WaitProcessInfo *[]ProcessInfo `json:"wait_process_info,omitempty"`
}

func (o MetaLockInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetaLockInfo struct{}"
	}

	return strings.Join([]string{"MetaLockInfo", string(data)}, " ")
}
