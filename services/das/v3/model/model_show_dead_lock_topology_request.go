package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeadLockTopologyRequest Request Object
type ShowDeadLockTopologyRequest struct {

	// 数据库用户ID。用户使用数据库账号与数据库建立的连接ID。
	ConnectionId string `json:"connection_id"`

	// 死锁的ID值
	DeadLockId string `json:"dead_lock_id"`
}

func (o ShowDeadLockTopologyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeadLockTopologyRequest struct{}"
	}

	return strings.Join([]string{"ShowDeadLockTopologyRequest", string(data)}, " ")
}
