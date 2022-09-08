package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExpandReplicationRequest struct {

	// 复制对的ID。
	ReplicationId string `json:"replication_id"`

	Body *ExtendReplicationRequestBody `json:"body,omitempty"`
}

func (o ExpandReplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandReplicationRequest struct{}"
	}

	return strings.Join([]string{"ExpandReplicationRequest", string(data)}, " ")
}
