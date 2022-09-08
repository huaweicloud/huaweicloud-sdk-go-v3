package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteReplicationRequest struct {

	// 复制对的ID。
	ReplicationId string `json:"replication_id"`

	Body *DeleteReplicationRequestBody `json:"body,omitempty"`
}

func (o DeleteReplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteReplicationRequest struct{}"
	}

	return strings.Join([]string{"DeleteReplicationRequest", string(data)}, " ")
}
