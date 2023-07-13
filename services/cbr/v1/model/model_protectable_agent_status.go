package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProtectableAgentStatus
type ProtectableAgentStatus struct {

	// agent无法连接的错误码
	Code *int32 `json:"code,omitempty"`

	// agent是否安装
	Installed *bool `json:"installed,omitempty"`

	// agent是否为老版本
	IsOld *bool `json:"is_old,omitempty"`

	// agent无法连接的错误信息
	Message *string `json:"message,omitempty"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// agent版本号
	Version *string `json:"version,omitempty"`
}

func (o ProtectableAgentStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectableAgentStatus struct{}"
	}

	return strings.Join([]string{"ProtectableAgentStatus", string(data)}, " ")
}
