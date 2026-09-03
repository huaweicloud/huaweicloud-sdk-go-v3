package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetSqlSwitchNewRequestBody 设置SQL开关请求体
type SetSqlSwitchNewRequestBody struct {

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 全量SQL开关
	FullSqlSwitchOn *bool `json:"full_sql_switch_on,omitempty"`

	// 全量SQL存储时长
	FullSqlRetentionHours *int64 `json:"full_sql_retention_hours,omitempty"`

	// 慢SQL开关
	SlowSqlSwitchOn *bool `json:"slow_sql_switch_on,omitempty"`

	// 慢SQL存储时长
	SlowSqlRetentionHours *int64 `json:"slow_sql_retention_hours,omitempty"`

	// 死锁开关
	DeadLockSwitchOn *bool `json:"dead_lock_switch_on,omitempty"`

	// 死锁存储时长
	DeadLockRetentionHours *int64 `json:"dead_lock_retention_hours,omitempty"`

	// 锁等待开关
	LockBlockingSwitchOn *bool `json:"lock_blocking_switch_on,omitempty"`

	// 锁等待存储时长
	LockBlockingRetentionHours *int64 `json:"lock_blocking_retention_hours,omitempty"`
}

func (o SetSqlSwitchNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSqlSwitchNewRequestBody struct{}"
	}

	return strings.Join([]string{"SetSqlSwitchNewRequestBody", string(data)}, " ")
}
