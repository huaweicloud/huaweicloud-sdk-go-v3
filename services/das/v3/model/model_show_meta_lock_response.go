package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMetaLockResponse Response Object
type ShowMetaLockResponse struct {

	// MDL锁等待信息
	Infos *[]MetaLockInfo `json:"infos,omitempty"`

	// MDL锁总数量
	Count *int32 `json:"count,omitempty"`

	// 等待锁的会话的数量
	WaitLockCount *int32 `json:"wait_lock_count,omitempty"`

	// 持有锁的会话的数量
	HoldLockCount *int32 `json:"hold_lock_count,omitempty"`

	// 等锁时间大于阈值的会话的数量
	TimeGreaterThanCount *int32 `json:"time_greater_than_count,omitempty"`

	// MDL锁等待时间阈值
	LockWaitThresholdSecond *int64 `json:"lock_wait_threshold_second,omitempty"`
	HttpStatusCode          int    `json:"-"`
}

func (o ShowMetaLockResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetaLockResponse struct{}"
	}

	return strings.Join([]string{"ShowMetaLockResponse", string(data)}, " ")
}
