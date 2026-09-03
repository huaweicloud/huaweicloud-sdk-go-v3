package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeadLockOriginDataRequest Request Object
type ShowDeadLockOriginDataRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 死锁ID
	DeadLockId string `json:"dead_lock_id"`

	// 开始时间戳 ms
	StartTime int64 `json:"start_time"`

	// 结束时间戳 ms
	EndTime int64 `json:"end_time"`
}

func (o ShowDeadLockOriginDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeadLockOriginDataRequest struct{}"
	}

	return strings.Join([]string{"ShowDeadLockOriginDataRequest", string(data)}, " ")
}
