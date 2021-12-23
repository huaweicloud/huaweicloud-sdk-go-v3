package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowReplicationRequest struct {
	// 复制对ID。

	ReplicationId string `json:"replication_id"`
}

func (o ShowReplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplicationRequest struct{}"
	}

	return strings.Join([]string{"ShowReplicationRequest", string(data)}, " ")
}
