package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeadLockRelationshipRequest Request Object
type ShowDeadLockRelationshipRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 死锁ID
	DeadLockId string `json:"dead_lock_id"`

	// 开始时间戳 ms
	StartTime int64 `json:"start_time"`

	// 结束时间戳 ms
	EndTime int64 `json:"end_time"`
}

func (o ShowDeadLockRelationshipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeadLockRelationshipRequest struct{}"
	}

	return strings.Join([]string{"ShowDeadLockRelationshipRequest", string(data)}, " ")
}
