package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetachProtectedInstanceReplicationRequest struct {

	// 保护实例的ID。
	ProtectedInstanceId string `json:"protected_instance_id" xml:"protected_instance_id"`

	// 复制对的ID。
	ReplicationId string `json:"replication_id" xml:"replication_id"`
}

func (o DetachProtectedInstanceReplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachProtectedInstanceReplicationRequest struct{}"
	}

	return strings.Join([]string{"DetachProtectedInstanceReplicationRequest", string(data)}, " ")
}
