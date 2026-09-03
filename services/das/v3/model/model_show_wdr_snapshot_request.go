package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWdrSnapshotRequest Request Object
type ShowWdrSnapshotRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *ShowWdrSnapshotRequestBody `json:"body,omitempty"`
}

func (o ShowWdrSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWdrSnapshotRequest struct{}"
	}

	return strings.Join([]string{"ShowWdrSnapshotRequest", string(data)}, " ")
}
