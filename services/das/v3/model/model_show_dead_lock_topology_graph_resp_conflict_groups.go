package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowDeadLockTopologyGraphRespConflictGroups struct {

	// 冲突组唯一标识
	GroupId string `json:"group_id"`

	// 等待锁节点唯一标识
	WaitingLockId string `json:"waiting_lock_id"`

	// 持有锁节点唯一标识
	GrantedLockId string `json:"granted_lock_id"`
}

func (o ShowDeadLockTopologyGraphRespConflictGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeadLockTopologyGraphRespConflictGroups struct{}"
	}

	return strings.Join([]string{"ShowDeadLockTopologyGraphRespConflictGroups", string(data)}, " ")
}
