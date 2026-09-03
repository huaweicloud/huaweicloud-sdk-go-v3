package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FullDeadLock 全量死锁信息
type FullDeadLock struct {

	// 发生时间（ms）
	HappenTime *int64 `json:"happen_time,omitempty"`

	// 死锁ID
	DeadLockId *string `json:"dead_lock_id,omitempty"`

	MysqlDeadLock *MySqlDeadLock `json:"mysql_dead_lock,omitempty"`
}

func (o FullDeadLock) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullDeadLock struct{}"
	}

	return strings.Join([]string{"FullDeadLock", string(data)}, " ")
}
