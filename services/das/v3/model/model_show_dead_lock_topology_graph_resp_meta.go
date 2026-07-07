package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeadLockTopologyGraphRespMeta 死锁概况
type ShowDeadLockTopologyGraphRespMeta struct {

	// 死锁唯一标识
	DeadLockId string `json:"dead_lock_id"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 死锁的发生时间，Unix 毫秒时间戳
	OccurTime int64 `json:"occur_time"`

	// 死锁的事务总数
	TotalTransactionsInCycle int32 `json:"total_transactions_in_cycle"`

	// 本次实际返回的事务数
	TotalTransactionsReturned int64 `json:"total_transactions_returned"`

	// 是否裁剪（>10 事务只返回 10 个）
	Truncated bool `json:"truncated"`
}

func (o ShowDeadLockTopologyGraphRespMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeadLockTopologyGraphRespMeta struct{}"
	}

	return strings.Join([]string{"ShowDeadLockTopologyGraphRespMeta", string(data)}, " ")
}
