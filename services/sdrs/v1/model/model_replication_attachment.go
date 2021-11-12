package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 复制对挂载信息数据结构
type ReplicationAttachment struct {
	// 该复制对挂载的保护实例ID。

	ProtectedInstance string `json:"protected_instance"`
	// 挂载点。

	Device string `json:"device"`
}

func (o ReplicationAttachment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplicationAttachment struct{}"
	}

	return strings.Join([]string{"ReplicationAttachment", string(data)}, " ")
}
