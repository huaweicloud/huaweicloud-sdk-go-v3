package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWaitingLocksSnapshotRequest Request Object
type ShowWaitingLocksSnapshotRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// InnoDB锁等待快照ID
	Id int32 `json:"id"`
}

func (o ShowWaitingLocksSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWaitingLocksSnapshotRequest struct{}"
	}

	return strings.Join([]string{"ShowWaitingLocksSnapshotRequest", string(data)}, " ")
}
