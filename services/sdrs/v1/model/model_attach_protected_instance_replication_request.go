package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AttachProtectedInstanceReplicationRequest struct {
	// 保护实例的ID。

	ProtectedInstanceId string `json:"protected_instance_id"`

	Body *ProtectedInstanceAttachReplicationRequestBody `json:"body,omitempty"`
}

func (o AttachProtectedInstanceReplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachProtectedInstanceReplicationRequest struct{}"
	}

	return strings.Join([]string{"AttachProtectedInstanceReplicationRequest", string(data)}, " ")
}
