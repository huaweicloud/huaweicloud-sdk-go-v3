package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FullDeadLockListRespFullDeadlockList struct {

	// 死锁事件发生的 Unix 时间戳（毫秒级）
	HappenTime *int64 `json:"happen_time,omitempty"`

	MysqlDeadlock *FullDeadLockListRespMysqlDeadlock `json:"mysql_deadlock,omitempty"`
}

func (o FullDeadLockListRespFullDeadlockList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullDeadLockListRespFullDeadlockList struct{}"
	}

	return strings.Join([]string{"FullDeadLockListRespFullDeadlockList", string(data)}, " ")
}
