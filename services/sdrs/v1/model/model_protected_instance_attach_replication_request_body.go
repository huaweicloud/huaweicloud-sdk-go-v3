package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProtectedInstanceAttachReplicationRequestBody 保护实例挂载复制对请求体
type ProtectedInstanceAttachReplicationRequestBody struct {
	ReplicationAttachment *ProtectedInstanceAttachReplicationRequestParams `json:"replicationAttachment"`
}

func (o ProtectedInstanceAttachReplicationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectedInstanceAttachReplicationRequestBody struct{}"
	}

	return strings.Join([]string{"ProtectedInstanceAttachReplicationRequestBody", string(data)}, " ")
}
