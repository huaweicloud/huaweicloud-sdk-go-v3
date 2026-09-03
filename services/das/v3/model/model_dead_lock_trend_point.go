package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeadLockTrendPoint 死锁趋势点
type DeadLockTrendPoint struct {

	// 发生时间
	OccurrenceTime *int64 `json:"occurrence_time,omitempty"`

	// 死锁总数
	TotalDeadlockCount *int64 `json:"total_deadlock_count,omitempty"`

	// keylock数量
	KeyDeadlockCount *int64 `json:"key_deadlock_count,omitempty"`

	// objectlock数量
	ObjectDeadlockCount *int64 `json:"object_deadlock_count,omitempty"`

	// ridlock数量
	RidDeadlockCount *int64 `json:"rid_deadlock_count,omitempty"`

	// pagelock数量
	PageDeadlockCount *int64 `json:"page_deadlock_count,omitempty"`

	// compilelock数量
	CompileDeadlockCount *int64 `json:"compile_deadlock_count,omitempty"`
}

func (o DeadLockTrendPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeadLockTrendPoint struct{}"
	}

	return strings.Join([]string{"DeadLockTrendPoint", string(data)}, " ")
}
