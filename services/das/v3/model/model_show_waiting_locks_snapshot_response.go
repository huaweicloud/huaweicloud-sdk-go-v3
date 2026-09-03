package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWaitingLocksSnapshotResponse Response Object
type ShowWaitingLocksSnapshotResponse struct {

	// InnoDB锁对应的事务的信息
	Trx *[]InnodbTrxInfo `json:"trx,omitempty"`

	// InnoDB锁等待信息
	LockWaitsInfos *[]interface{} `json:"lock_waits_infos,omitempty"`

	// 等待锁的会话的数量
	WaitLockCount *int32 `json:"wait_lock_count,omitempty"`

	// 持有锁的会话的数量
	HoldLockCount *int32 `json:"hold_lock_count,omitempty"`

	// 等锁时间大于阈值的会话的数量
	TimeGreaterThanCount *int32 `json:"time_greater_than_count,omitempty"`

	// InnoDB锁等待时间阈值
	LockWaitThresholdSecond *int64 `json:"lock_wait_threshold_second,omitempty"`
	HttpStatusCode          int    `json:"-"`
}

func (o ShowWaitingLocksSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWaitingLocksSnapshotResponse struct{}"
	}

	return strings.Join([]string{"ShowWaitingLocksSnapshotResponse", string(data)}, " ")
}
