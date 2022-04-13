package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 备份时跳过的资源
type CheckpointCreateSkippedResource struct {
	// 资源ID

	Id *string `json:"id,omitempty"`
	// 资源类型

	Type *string `json:"type,omitempty"`
	// 资源名称

	Name *string `json:"name,omitempty"`
	// 请参见[错误码](ErrorCode.xml)。

	Code *string `json:"code,omitempty"`
	// 跳过原因，例如：该资源正在备份中。

	Reason *string `json:"reason,omitempty"`
}

func (o CheckpointCreateSkippedResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointCreateSkippedResource struct{}"
	}

	return strings.Join([]string{"CheckpointCreateSkippedResource", string(data)}, " ")
}
